package main

import (
	"log"
	"net/http"

	"github.com/souvikhaldar/issue-warden/issue"
	"github.com/souvikhaldar/issue-warden/user"
)

func main() {
	http.HandleFunc("/login", user.CheckLogin)
	http.HandleFunc("/signup", user.Signup)
	http.HandleFunc("/verify", user.Verify)
	http.HandleFunc("/issue", issue.AddIssue())
	log.Fatal(http.ListenAndServe(":8192", nil))

}
