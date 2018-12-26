package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/souvikhaldar/issue-warden/data"
)

// Signup is used for creating a new user
func Signup(w http.ResponseWriter, r *http.Request) {
	var userid int
	var userDetails User
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&userDetails)
	if er != nil {
		oops := fmt.Sprint("Unable to decode user json ", er)
		http.Error(w, oops, http.StatusInternalServerError)
		log.Fatal(oops)
		return
	}
	e := data.DbConn.QueryRow(data.InsertUser, userDetails.Email, userDetails.Username, userDetails.FirstName, userDetails.LastName, userDetails.Password, userDetails.AccessToken).Scan(&userid)
	if e != nil {
		oops := fmt.Sprint("Error in inserting to user table ", e)
		log.Fatal(oops)
		http.Error(w, oops, http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully added user with userid: %d", userid)
}
