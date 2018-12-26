package main

import (
	"fmt"

	"github.com/souvikhaldar/issue-warden/user"
)

func main() {
	var email, password string
	var decision string
	fmt.Println("Do you wish to login or sign up? (login/signup)")
	fmt.Scanf("%s", &decision)
	if decision == "login" {
		fmt.Println("Enter email: ")
		fmt.Scanf("%s", &email)
		fmt.Println("Enter password: ")
		fmt.Scanf("%s", &password)
		user.CheckLogin(email, password)
	}

}
