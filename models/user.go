package models

import (
	"log"

	d "github.com/tavvfiq/cafe-rest-api/database"
	tb "github.com/tavvfiq/cafe-rest-api/database/mysql"
)

// RegisterUser Registing user to the database
func RegisterUser(b *tb.RegistReq) (tb.User, error) {
	userTable := tb.User{}
	stmt, err := d.DB.Prepare("INSERT INTO users SET ?")
	if err != nil {
		log.Fatal(err)
		return tb.User{}, err
	}

	_, err = stmt.Exec(b)
	if err != nil {
		log.Fatal(err)
		return tb.User{}, err
	}
	stmt.Close()
	rows, err := d.DB.Query("SELECT id, first_name, last_name, phone_number, profile_image, level_id FROM users WHERE users.email=?;", "taufiqwidinugroho@gmail.com")
	if err != nil {
		log.Fatal(err)
		return tb.User{}, err
	}
	for rows.Next() {
		if err := rows.Scan(&userTable.FirstName, &userTable.LastName, &userTable.Email, &userTable.ProfileImage, &userTable.PhoneNumber, &userTable.LevelID); err != nil {
			log.Fatal(err)
			return tb.User{}, err
		}
	}
	return userTable, nil
}
