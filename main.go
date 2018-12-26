package main

import (
	"log"
	"net/http"

	"github.com/souvikhaldar/issue-warden/user"
)

func main() {
	http.HandleFunc("/login", user.CheckLogin)
	http.HandleFunc("/signup", user.Signup)
	log.Fatal(http.ListenAndServe(":8192", nil))

}
