package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/souvikhaldar/gorand"
	"github.com/souvikhaldar/issue-warden/data"
)

// Signup is used for creating a new user
func Signup(w http.ResponseWriter, r *http.Request) {
	var userid int
	var upsert string
	var userDetails User
	decoder := json.NewDecoder(r.Body)
	er := decoder.Decode(&userDetails)
	if er != nil {
		oops := fmt.Sprint("Unable to decode user json ", er)
		fmt.Println(oops)
		http.Error(w, oops, http.StatusInternalServerError)
		return
	}
	fmt.Printf("User details recieved is: %+v", userDetails)
	e := data.DbConn.QueryRow(data.InsertUser, userDetails.Email, userDetails.Username, userDetails.FirstName, userDetails.LastName, userDetails.Password).Scan(&userid, &upsert)
	if e != nil {
		oops := fmt.Sprint("Error in inserting to user table ", e)
		fmt.Println(oops)
		http.Error(w, oops, http.StatusInternalServerError)
		return
	}
	if upsert == "inserted" {
		fmt.Fprintf(w, "Successfully added user with userid: %d \n", userid)
		otp := gorand.RandStrWithCharset(6, "123456")
		useridString := strconv.Itoa(userid)
		if _, e := RedisClient.Set(useridString, otp, 300*time.Second).Result(); e != nil {
			oops := fmt.Sprintln("Unable to set otp ", e)
			fmt.Println(oops)
			http.Error(w, oops, http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Check email and verify the otp against the userid:%d and otp:%s ", userid, otp)
		// send mail to given email with otp and timer of 5 min
		if e := Mailconfig.SendMail([]string{userDetails.Email}, "OTP for issue-warden", fmt.Sprintf("The OTP is %s. \n Regards. \n Team Issue-Warden", otp)); e != nil {
			fmt.Printf("Error in sending email %s ", e)
		}

	} else if upsert == "updated" {
		oops := fmt.Sprintf("User with userid %d already exists, please log in instead ", userid)
		fmt.Println(oops)
		http.Error(w, oops, http.StatusBadRequest)
	}
}
