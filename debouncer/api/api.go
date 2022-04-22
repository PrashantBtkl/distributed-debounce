package api

import (
    "net/http"
    "strconv"
    "time"

    "github.com/PrashantBtkl/distributed-debounce/debouncer/model"
    "github.com/PrashantBtkl/distributed-debounce/debouncer/services/rabbitmq"
    "github.com/PrashantBtkl/distributed-debounce/debouncer/services/store"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
)

const (
    host    = "127.0.0.1"
    port    = "1234"
    service = "debouncer"
)

type HTTPListener struct {
    host string
    port string
    srv  *http.Server
    db   *store.PGStore
    mq   *rabbitmq.RabbitMQ
}

func (l *HTTPListener) homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("<html>Pong</html>"))
}

func (l *HTTPListener) debounceHandler(w http.ResponseWriter, r *http.Request) {
    v := mux.Vars(r)
    //1.update/create entry on debounce table
    userID, err := strconv.Atoi(v["id"])
    if err != nil {
        log.WithFields(log.Fields{
            "service": service,
            "err":     err.Error(),
        }).Error("failed to convert input to int")

    } else {
        l.db.UpdateBuffer(userID, 10000)
    }
    var rmqBody []byte
    rmqBody = []byte(strconv.Itoa(userID))
    //2. send message to rabbitmq
    l.mq.PublishWithDelay("delayed", rmqBody, 10000)
    if err != nil {
        log.WithFields(log.Fields{
            "service": service,
            "err":     err.Error(),
        }).Error("failed to publish message to rabbitmq")
    }

}

func NewHTTPListener(host string, port string) error {

    db, err := store.NewPGStore(store.PGDSN("127.0.0.1", "postgres", "password", "debounce", "5432"))
    if err != nil {
        log.WithFields(log.Fields{
            "service": service,
            "err":     err.Error(),
        }).Error("Failed to create db")
        return err
    }
    var amqp = model.AMQP{URL: "amqp://guest:guest@127.0.0.1:5672/", Exchange: "amq.direct"}

    rmq, err := rabbitmq.InitRabbitMQ(amqp, db)
    if err != nil {
        log.Fatalf("run: failed to init rabbitmq: %v", err)
    }

    ret := &HTTPListener{
        host: host,
        port: port,
        db:   db,
        mq:   rmq,
    }

    r := mux.NewRouter()
    api_version := "api/v1"
    r.HandleFunc("/ping", ret.homeHandler)
    r.HandleFunc("/"+api_version+"/{id:[0-9]+}", ret.debounceHandler).Methods("GET")

    srv := &http.Server{
        Handler: r,
        Addr:    host + ":" + port,
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 60 * time.Second,
        ReadTimeout:  60 * time.Second,
    }

    ret.srv = srv
    return nil
}
