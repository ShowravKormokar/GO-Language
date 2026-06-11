package models

import (
	"time"
)

type BlacklistedToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	JTI       string    `gorm:"type:varchar(36);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null;index"`
	CreatedAt time.Time
}
