package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
