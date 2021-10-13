package models

type User struct {
	//ID       *string `json : "id, omitempty" bson: "_id, omitempty"`
	Name     *string `json:"name"`
	Price    *int64  `json:"price"`
	Quantity *int64  `json:"quantity"`
}
