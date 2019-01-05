package gomail

import (
	"errors"
	"fmt"
	"net/smtp"
)

type Config struct {
	From     string
	Password string
}

func New(from string, password string) (error, *Config) {
	if from == "" {
		return errors.New("To and from field can't be empty"), &Config{}
	}
	return nil, &Config{
		From:     from,
		Password: password,
	}
}

func (c *Config) CreateMsg(To []string, Subject, Body string) []string {
	msgString := make([]string, len(To))
	for i, to := range To {
		msg := fmt.Sprintf("To: %s \r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", to, Subject, Body)
		msgString[i] = msg
	}
	return msgString
}

func (c *Config) SendMail(To []string, subject, body string) error {
	auth := smtp.PlainAuth("", c.From, c.Password, "smtp.gmail.com")
	for _, to := range To {
		msg := fmt.Sprintf("To: %s \r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", to, subject, body)
		err := smtp.SendMail("smtp.gmail.com:587", auth, c.From, []string{to}, []byte(msg))
		if err != nil {
			return err
		}
	}
	return nil
}
