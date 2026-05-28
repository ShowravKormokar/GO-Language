package models

import "time"

type Contact struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
