package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/souvikhaldar/issue-warden/data"
)

// CheckLogin checks whether the email and password is valid or not
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"][0]
	var pass string
	e := data.DbConn.QueryRow(data.CheckPassword, email).Scan(&pass)
	if e != nil {
		oops := fmt.Sprint("Unable to fetch password ", e)
		log.Fatal(oops)
		http.Error(w, oops, http.StatusInternalServerError)
		return
	}
	if pass == password {
		fmt.Fprintf(w, "Successfully logged in")
	} else {
		fmt.Fprintf(w, "Invalid username or password")
	}
}
