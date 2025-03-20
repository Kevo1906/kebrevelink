package models

import (
	"time"

	"github.com/google/uuid"	
)

type URL struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ShortUrl  string         `gorm:"unique"`
	Original  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	ExpiresAt *time.Time
	UserID    uuid.UUID      `gorm:"type:uuid;index"`
	User      User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}