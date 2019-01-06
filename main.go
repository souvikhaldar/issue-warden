package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/souvikhaldar/issue-warden/issue"
	"github.com/souvikhaldar/issue-warden/user"
)

func main() {
	durPtr := flag.Int("duration", 86400, "The interval of mailing issues in seconds")
	flag.Parse()
	go issue.SendIssueMails(time.Duration(*durPtr))
	http.HandleFunc("/login", user.CheckLogin)
	http.HandleFunc("/signup", user.Signup)
	http.HandleFunc("/verify", user.Verify)
	http.HandleFunc("/issue", issue.AddIssue())
	log.Fatal(http.ListenAndServe(":8192", nil))

}
