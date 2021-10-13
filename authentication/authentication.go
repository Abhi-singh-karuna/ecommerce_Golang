package authentication

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var jwtWare = "gosecretkey"

var u = models.User{
	Email:    "email",
	Password: "password",
}

//var client *mongo.Client

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(jwtWare)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func SignUp(c *fiber.Ctx) error {

	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	user.Password = getHash([]byte(user.Password))
	collection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, user)

	if result != nil {
		fmt.Println("User Registered")
	}
	return c.JSON(user)
}

func CreateToken(userMail string) (string, error) {
	var err error
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_mail"] = userMail
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("gosecretkey"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Login(c *fiber.Ctx) error {

	user := new(models.User)
	var dbUser models.User
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	collection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if err != nil {
		return err
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		fmt.Println("Wrong Password")
	}
	token, err := CreateToken(u.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity)
		return err
	}
	c.JSON(token)
	return nil
}
