package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	URLs      []URL          `gorm:"foreignKey:UserID"`
}