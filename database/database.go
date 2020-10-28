package database

import (
	"database/sql"
	"log"

	// import driver
	_ "github.com/go-sql-driver/mysql"
)

// DB exported database
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:8000)/backend_expressjs")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mysql database at port: 8000")
	}
}
