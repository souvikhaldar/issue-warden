package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func init() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = ""
		dbname   = "issuewarden"
	)
	var er error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	DbConn, er = sql.Open("postgres", psqlInfo)
	if er != nil {
		fmt.Println("Unable to connect to the database ", er)
	}

}
