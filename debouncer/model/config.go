package model

// AMQP is the amqp configuration
type AMQP struct {
	URL      string `default:"amqp://guest:guest@127.0.0.1:5672/guest"`
	Exchange string `default:"amq.direct"`
}

type DB struct {
	PGUser     string `default:"postgres"`
	PGDB       string `default:"debounce"`
	PGHost     string `default:"127.0.0.1"`
	PGPort     string `default:"5432"`
	PGPassword string `default:"password"`
}

// Config is the application configuration
type Config struct {
	AppName    string `json:"app-name" default:"rabbitmq"`
	AppVersion string `json:"app-version" required:"true"`

	AMQP AMQP
	DB   DB
}
