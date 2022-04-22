package model

import (
	uuid "github.com/satori/go.uuid"
)

type DebounceBuffer struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID         int       `gorm:"column:user_id"`
	DebounceBuffer int64     `gorm:"column:debounce_buffer"`
}
