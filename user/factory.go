package user

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/souvikhaldar/gomail"
)

// User contains the details about the customer
type User struct {
	Userid      int
	Email       string
	Username    string
	FirstName   string
	LastName    string
	Password    string
	AccessToken string
}

var RedisClient *redis.Client
var Mailconfig *gomail.Config

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var e error
	e, Mailconfig = gomail.New("joeymartin367@gmail.com", "Hu$tl3r11")
	if e != nil {
		fmt.Print(fmt.Errorf("Error in creating config %v", e))
	}

}

func New(email string, username string, firstname string, lastname string, password string, accesstoken string) *User {
	return &User{
		Email:       email,
		Username:    username,
		FirstName:   firstname,
		LastName:    lastname,
		Password:    password,
		AccessToken: accesstoken,
	}
}
