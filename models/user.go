package models

import (
	"log"

	"context"

	d "github.com/tavvfiq/cafe-rest-api/database"
	tb "github.com/tavvfiq/cafe-rest-api/database/mysql"
)

// RegisterUser Registering user to the database
func RegisterUser(c context.Context, b *tb.RegistReq) (tb.User, error) {
	userTable := tb.User{}

	stmt, err := d.DB.Prepare("INSERT INTO users SET ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(b)
	if err != nil {
		panic(err)
	}
	rows, err := d.DB.QueryContext(c, "SELECT id, first_name, last_name, phone_number, profile_image, level_id FROM users WHERE users.email=?;", b.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&userTable.FirstName, &userTable.LastName, &userTable.Email, &userTable.ProfileImage, &userTable.PhoneNumber, &userTable.LevelID); err != nil {
			log.Fatal(err)
			return tb.User{}, err
		}
	}
	return userTable, nil
}
