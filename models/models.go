package models

//--------------Created User struct for user registration-------------//
type User struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

//--------------Created Response struct to return the success/failure on signup-------------//
type Response struct {
	Message string `json:"message"`
}