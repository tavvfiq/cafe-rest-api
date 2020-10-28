package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	// import driver
	_ "github.com/go-sql-driver/mysql"
)

// DB exported database
var DB *sql.DB

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// get variable from .env
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	// connect to database
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName))

	if err != nil {
		log.Fatal(err)
	} else {
		if err := DB.PingContext(ctx); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Connected to mysql database")
		}
	}
}
