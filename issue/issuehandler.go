package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/souvikhaldar/issue-warden/data"
	"github.com/souvikhaldar/issue-warden/user"
)

func AddIssue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userid := r.URL.Query()["userid"]
		token := r.URL.Query()["token"]
		action := r.URL.Query()["action"]
		retrievedToken, err := user.RedisClient.Get(userid[0]).Result()
		if err != nil {
			oops := fmt.Sprintf("Unable to retrive token from redis db ", err)
			fmt.Println(oops)
			http.Error(w, oops, http.StatusInternalServerError)
		}
		fmt.Println("Retrieved token and recieved token are: ", retrievedToken, token)
		if retrievedToken == token[0] {
			fmt.Println("User authenticated")
			if action[0] == "insert" {
				var is Issue
				decoder := json.NewDecoder(r.Body)
				er := decoder.Decode(&is)
				if er != nil {
					oops := fmt.Sprint("Unable to decode issue json ", er)
					fmt.Println(oops)
					http.Error(w, oops, http.StatusInternalServerError)
					return
				}
				var issueID int
				e := data.DbConn.QueryRow(data.InsertIssue, is.Title, is.Description, is.AssignedTo, userid[0], false).Scan(&issueID)
				if e != nil {
					oops := fmt.Sprint("Error in inserting to issue table ", e)
					fmt.Println(oops)
					http.Error(w, oops, http.StatusInternalServerError)
					return
				}
				go sendIssueMail(is, 720)
				fmt.Fprintf(w, "Issue added %d", issueID)
			} else if action[0] == "delete" {
				var issueid string
				decoder := json.NewDecoder(r.Body)
				er := decoder.Decode(&issueid)
				if er != nil {
					oops := fmt.Sprint("Unable to decode issue json ", er)
					fmt.Println(oops)
					http.Error(w, oops, http.StatusInternalServerError)
					return
				}
				issueID, _ := strconv.Atoi(issueid)
				if _, e := data.DbConn.Exec("delete from issue where id=$1", issueID); e != nil {
					ops := fmt.Sprintf("Unable to delete %s because of:%s", issueid, e)
					fmt.Println(ops)
					http.Error(w, ops, 500)
					return
				}
				fmt.Fprintf(w, "Issue deleted")
			} else if action[0] == "update" {
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
				e := data.DbConn.QueryRow(data.UpdateIssue, is.Title, is.Description, is.AssignedTo, is.CreatedBy, is.Status, is.Id).Scan(&issueID)
				if e != nil {
					oops := fmt.Sprint("Error in updating to issue table ", e)
					fmt.Println(oops)
					http.Error(w, oops, http.StatusInternalServerError)
					return
				}
				go sendIssueMail(is, 1)
				fmt.Fprint(w, "Issue Updated")
			}
		} else {
			fmt.Println("Token mismatch, permission denied")
			http.Error(w, "Token mismatch, permission denied", 403)
		}
	}

}

func sendIssueMail(is Issue, dur time.Duration) {
	time.Sleep(dur * time.Second)
	if e := user.Mailconfig.SendMail([]string{is.AssignedTo}, fmt.Sprintf("Task: %s ", is.Title), fmt.Sprintf("You have been assigned task: %s \nCurrent Status: %t \nBy: %s", is.Description, is.Status, is.CreatedBy)); e != nil {
		err := fmt.Sprintf("Unable to send email for issue %s", e)
		fmt.Println(err)
		return
	}
}

func SendIssueMails(dur time.Duration) {
	var is Issue
	for {
		time.Sleep(dur * time.Second)
		rows, err := data.DbConn.Query("select * from issue")
		if err != nil {
			ops := fmt.Errorf("Error in fetching issues %s", err)
			fmt.Println(ops)
		}
		defer rows.Close()
		for rows.Next() {
			if e := rows.Scan(&is.Id, &is.Title, &is.Description, &is.AssignedTo, &is.CreatedBy, &is.Status); e != nil {
				ops := fmt.Errorf("Error in looping through issues %s", e)
				fmt.Println(ops)
			}
			if is.Id != 0 {
				fmt.Printf("Remainder mail sent to : %s\n", is.AssignedTo)
				sendIssueMail(is, 1)
			} else {
				fmt.Println("No Issues found")
			}
		}
	}
}
