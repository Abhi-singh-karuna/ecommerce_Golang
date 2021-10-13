package models

import (
	"time"
)

//
type Product struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        *string   `json:"title"`
	Price       *int      `json:"price"`
	Quantity    *int      `json:"Quantity"`
	Description *string   `json:"Description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductCart struct {
	Name     *string `json:"title"`
	Price    *int    `json:"price"`
	Quantity *int    `json:"Quantity"`
}

type Total struct {
	ID    *string `json:"id,omitempty" bson:"_id,omitempty"`
	Total *int    `json:"total"`
}

type Add struct {
	ID        *string `json:"id,omitempty" bson:"_id,omitempty"`
	Full_Name *string `json:"full_Name"`
	Address_a *string `json:"address_a"`
	Address_b *string `json:"address_b"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	PinCode   *int    `json:"pincode"`
}
