package user

// User contains the details about the customer
type User struct {
	Userid      int
	Email       string
	Username    string
	FirstName   string
	LastName    string
	Password    string
	AccessToken string
}

func New(email string, username string, firstname string, lastname string, password string, accesstoken string) *User {
	return &User{
		Email:       email,
		Username:    username,
		FirstName:   firstname,
		LastName:    lastname,
		Password:    password,
		AccessToken: accesstoken,
	}
}
