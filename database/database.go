package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	// import driver
	_ "github.com/go-sql-driver/mysql"
)

// DB exported database
var DB *sql.DB

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	DB, err = sql.Open("mysql", "root:@tcp(localhost)/backend_expressjs")
	if err != nil {
		log.Fatal(err)
	} else {
		if err := DB.PingContext(ctx); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Connected to mysql database at port: 8000")
		}
	}
}
