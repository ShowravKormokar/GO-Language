package model

import "time"

type Product struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"name"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64   `json:"price" bson:"price"`
	Qunatity    int64     `json:"quantity" bson:"quantity"`
	Status      bool      `json:"status" bson:"status"`
	Created_At  time.Time `json:"created_at" bson:"created_at"`
	Updated_At  time.Time `json:"updated_at" bson:"updated_at"`
}
