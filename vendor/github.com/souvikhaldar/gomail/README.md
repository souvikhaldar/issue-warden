# gomail
A simple golang package to send email using gmail

## Usage

```
package main

import (
	"fmt"

	"github.com/souvikhaldar/gomail"
)

func main() {
	e, config := gomail.New("<from>", []string{"<to>"}, "<subject>", "<body>", "<password>")
	if e != nil {
		fmt.Print(fmt.Errorf("Error in creating config %v", e))
	}
	if e := config.SendMail(); e != nil {
		fmt.Print(fmt.Errorf("Error in sending mail %v", e))
	}
}
```

## Note
Gmail need to be allowed access to unsafe app.

https://serverfault.com/questions/635139/how-to-fix-send-mail-authorization-failed-534-5-7-14

