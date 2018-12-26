package user

import (
	"fmt"
	"net/http"

	"github.com/souvikhaldar/issue-warden/data"
)

// CheckLogin checks whether the email and password is valid or not
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query()["email"]
	password := r.URL.Query()["password"]
	fmt.Println("Email and password is ", string(email[0]), string(password[0]))
	var pass string
	e := data.DbConn.QueryRow(data.CheckPassword, string(email[0])).Scan(&pass)
	if e != nil {
		oops := fmt.Sprint("Unable to fetch password ", e)
		fmt.Println(oops)
		http.Error(w, oops, http.StatusInternalServerError)
		return
	}
	if pass == password[0] {
		fmt.Fprintf(w, "Successfully logged in")
	} else {
		fmt.Fprintf(w, "Invalid username or password")
	}
}
