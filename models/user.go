package models

import (
	"agt2020/event-booking/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) SaveUser() (int64, error) {
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
	err = stmt.QueryRow(&u.Email, &u.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
