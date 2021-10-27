package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
	Quantity    int                `json:"Quantity"`
	Description string             `json:"Description"`
	Image       string             `json:"image"`
	//CreatedAt   time.Time `json:"createdAt"`
	//UpdatedAt   time.Time `json:"updatedAt"`
}

type Productbyid struct {
}
