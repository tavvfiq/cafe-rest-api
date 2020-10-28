package models

import (
	"context"

	"database/sql"

	"errors"

	d "github.com/tavvfiq/cafe-rest-api/database"
	tb "github.com/tavvfiq/cafe-rest-api/database/mysql"
	bcrypt "golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

// RegisterUser Registering user to the database
func RegisterUser(c context.Context, b *tb.RegistReq) (tb.User, error) {
	userTable := tb.User{}

	// hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), salt)

	// hashing error
	if err != nil {
		panic(err)
	}

	stmt, err := d.DB.Prepare("INSERT INTO users(first_name, last_name, email, password, level_id) VALUES(?,?,?,?,?);")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(b.FirstName, b.LastName, b.Email, hashedPassword, b.LevelID)
	if err != nil {
		return tb.User{}, errors.New("Email already registered")
	}

	var ProfileImage sql.NullString
	var PhoneNumber sql.NullString

	err = d.DB.QueryRowContext(c, "SELECT id, first_name, last_name, email, phone_number, profile_image, level_id FROM users WHERE users.email=?;", b.Email).Scan(&userTable.ID, &userTable.FirstName, &userTable.LastName, &userTable.Email, &ProfileImage, &PhoneNumber, &userTable.LevelID)
	if err != nil {
		panic(err)
	}

	if ProfileImage.Valid {
		userTable.ProfileImage = ProfileImage.String
	}
	if PhoneNumber.Valid {
		userTable.PhoneNumber = PhoneNumber.String
	}

	return userTable, nil
}

//LoginUser check user credential for authentication
func LoginUser(c context.Context, b *tb.LoginReq) (tb.User, error) {
	var password string
	var ProfileImage sql.NullString
	var PhoneNumber sql.NullString
	userTable := tb.User{}
	err := d.DB.QueryRowContext(c, "SELECT id, first_name, last_name, email, phone_number, profile_image, password, level_id FROM users WHERE email=?;", b.Email).Scan(&userTable.ID, &userTable.FirstName, &userTable.LastName, &userTable.Email, &PhoneNumber, &ProfileImage, &password, &userTable.LevelID)

	if err != nil {
		return tb.User{}, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(b.Password))
	if err != nil {
		return tb.User{}, errors.New("Wrong password")
	}

	return userTable, nil
}
