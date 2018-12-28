package user

import (
	"fmt"

	"github.com/go-redis/redis"
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

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
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
