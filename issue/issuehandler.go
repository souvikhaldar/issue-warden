package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/souvikhaldar/issue-warden/data"
	"github.com/souvikhaldar/issue-warden/user"
)

func AddIssue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query()["userid"]
		token := r.URL.Query()["token"]
		retrievedToken, err := user.RedisClient.Get(userid[0]).Result()
		if err != nil {
			oops := fmt.Sprintf("Unable to retrive token from redis db ", err)
			fmt.Println(oops)
			http.Error(w, oops, http.StatusInternalServerError)
		}
		fmt.Println("Retrieved token and recieved token are: ", retrievedToken, token)
		if retrievedToken == token[0] {
			fmt.Println("User authenticated")
			var is Issue
			decoder := json.NewDecoder(r.Body)
			er := decoder.Decode(&is)
			if er != nil {
				oops := fmt.Sprint("Unable to decode issue json ", er)
				fmt.Println(oops)
				http.Error(w, oops, http.StatusInternalServerError)
				return
			}
			fmt.Printf("Issue details recieved is: %+v", is)
			var issueID int
			e := data.DbConn.QueryRow(data.InsertIssue, is.Title, is.Description, is.AssignedTo, userid[0], false).Scan(&issueID)
			if e != nil {
				oops := fmt.Sprint("Error in inserting to issue table ", e)
				fmt.Println(oops)
				http.Error(w, oops, http.StatusInternalServerError)
				return
			}
			go sendIssueMail(is, w)

			fmt.Fprint(w, "Issue added")
		} else {
			fmt.Println("Token mismatch, permission denied")
			http.Error(w, "Token mismatch, permission denied", 403)
		}
	}
}

func sendIssueMail(is Issue, w http.ResponseWriter) {
	time.Sleep(12 * time.Minute)
	if e := user.Mailconfig.SendMail([]string{is.AssignedTo}, fmt.Sprintf("You have been assigned task: %s by %s", is.Title, is.CreatedBy), is.Description); e != nil {
		err := fmt.Sprintf("Unable to send email for issue %s", e)
		fmt.Println(err)
		http.Error(w, err, 500)
	}
}
