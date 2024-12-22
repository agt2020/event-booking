package models

import (
	"agt2020/event-booking/db"
	"agt2020/event-booking/utils"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding= "required"`
	Password string `binding= "required"`
}

func (u User) SaveUser() (int64, error) {
	query := `
	INSERT INTO public.users (email, password)
	VALUES ($1, $2) RETURNING id
	`
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var userID int64
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u User) Auth() error {
	query := "SELECT password FROM public.users WHERE email=$1"
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(u.Email).Scan(&hashedPassword)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	passwordIsValid := utils.CheckPassword(u.Password, hashedPassword)
	if passwordIsValid {
		return nil
	} else {
		return fmt.Errorf("password is invalid")
	}
}
