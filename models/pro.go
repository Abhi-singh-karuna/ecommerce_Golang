package models

import (
	"time"
)


type Product struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        *string   `json:"title"`
	Price       *float64  `json:"price"`
	Quantity    *int      `json:"Quantity"`
	Description *string   `json:"Description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
