package user

import (
	"fmt"
	"net/http"

	"github.com/souvikhaldar/gorand"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	userid := r.URL.Query()["userid"]
	otp := r.URL.Query()["otp"]
	fmt.Println("Userid and otp are: ", userid[0], otp[0])
	retrievedOtp, err := RedisClient.Get(userid[0]).Result()
	if err != nil {
		oops := fmt.Sprintf("Unable to retrive otp from redis db %s", err)
		fmt.Println(oops)
		http.Error(w, oops, http.StatusInternalServerError)
	}
	if otp[0] == retrievedOtp {
		yay := fmt.Sprintf("User %s is now verified ", userid[0])
		fmt.Println(yay)
		fmt.Fprintf(w, yay)
		// setting access token to the user
		token := gorand.RandStr(10)
		if _, e := RedisClient.Set(userid[0], token, 0).Result(); e != nil {
			oops := fmt.Sprintf("Unable to set token for the user ", e)
			fmt.Errorf("%s", oops)
			http.Error(w, oops, http.StatusInternalServerError)
		}
	} else {
		fmt.Printf("Couldn't verify the OTP")
		http.Error(w, "OTP invalid/incorrect ", http.StatusUnauthorized)
	}
}
