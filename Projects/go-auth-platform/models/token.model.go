package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	TokenHash string    `gorm:"type:varchar(64);uniqueIndex;not null"`

	ExpiresAt time.Time `gorm:"not null;index"`
	Revoked   bool      `gorm:"default:false"`

	CreatedAt time.Time
}

type BlacklistedToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	JTI       string    `gorm:"type:varchar(36);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null;index"`
	CreatedAt time.Time
}
