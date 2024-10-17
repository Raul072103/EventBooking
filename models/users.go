package models

import (
	"EventBooking/db"
	"EventBooking/utils"
	"database/sql"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func Save(u *User) error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = errors.Join(err, stmt.Close())
	}(stmt)

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	u.Id, err = result.LastInsertId()
	return err
}

func ValidateCredentials(u *User) error {
	query := `SELECT id, password FROM users WHERE email=?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var userId int64

	err := row.Scan(&userId, &retrievedPassword)
	if err != nil {
		return errors.New("credentials failed")
	}

	passwordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("credentials failed")
	}

	u.Id = userId
	return nil
}
