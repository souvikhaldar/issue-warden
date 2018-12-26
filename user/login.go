package user

import (
	"fmt"

	"github.com/souvikhaldar/issue-warden/data"
)

func CheckLogin(email, password string) bool {

}
func SignUp(userDetails *User) int {
	var userid int
	e := data.DbConn.QueryRow(data.InsertUser, userDetails.Email, userDetails.Username, userDetails.FirstName, userDetails.LastName, userDetails.Password, userDetails.AccessToken).Scan(&userid)
	if e != nil {
		fmt.Println("Error in inserting to user table ", e)
		return 0
	}
	return userid
}
