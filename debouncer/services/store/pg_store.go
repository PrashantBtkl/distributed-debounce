package store

import (
	"fmt"
	"github.com/PrashantBtkl/distributed-debounce/debouncer/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"time"

	"gorm.io/gorm"
)

type PGStore struct {
	db *gorm.DB
}

func PGDSN(dbhost, dbusername, dbpassword, dbname, dbport string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbhost, dbusername, dbpassword, dbname, dbport)
}

func NewPGStore(dsn string) (*PGStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open DB: %v", err)
		return nil, err
	}
	return &PGStore{
		db: db,
	}, nil
}

func (db *PGStore) UpdateBuffer(user int, bufferDuration int64) {
	debounce := &model.DebounceBuffer{UserID: user, DebounceBuffer: time.Now().Unix() + bufferDuration}
	if db.db.Model(&model.DebounceBuffer{}).Where("user_id = ?", user).Updates(debounce).RowsAffected == 0 {
		db.db.Create(debounce)
	}

}

func (db *PGStore) CheckBuffer(user int) (*model.DebounceBuffer, error) {
	debounce := &model.DebounceBuffer{UserID: user}
	err := db.db.Find(debounce).Error
	if err != nil {
		return nil, err
	}
	return debounce, nil
}
