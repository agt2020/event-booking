package models

import (
	"agt2020/event-booking/db"
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding="required"`
	Description string `binding="required"`
	Location    string `binding="required"`
	DateTime    time.Time
	UserID      int64
}

var events = []Event{}

func (e Event) Save() (int64, error) {
	query := `
	INSERT INTO public.events (name, description, location, dateTime, user_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	var lastID int64

	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&lastID)
	if err != nil {
		return -1, fmt.Errorf("failed to execute statement: %w", err)
	}
	return lastID, nil
}

func (e Event) Update() (int, error) {
	query := `
	UPDATE public.events SET name=$1, description=$2, location=$3, dateTime=$4, user_id=$5
	WHERE id=$6 RETURNING id
	`
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastID int
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID).Scan(&lastID)
	if err != nil || lastID == 0 {
		return 0, fmt.Errorf("failed to update event: %w", err)
	}

	return lastID, nil
}

func DeleteEvent(id string) error {
	query := "DELETE FROM public.events WHERE id=$1"
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	isDeleted, err := result.RowsAffected()
	if isDeleted == 1 {
		return nil
	} else {
		return err
	}
}

func GetAllEvents() (*[]Event, error) {
	events = []Event{}

	query := "SELECT * FROM public.events"
	DB := db.Initdb()
	rows, err := db.RunQuery(DB, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event

		// Scan the columns into variables
		err := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Description,
			&e.Location,
			&e.DateTime,
			&e.UserID,
		)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
		}

		events = append(events, e)
	}
	return &events, err
}

func GetEvent(id string) (*Event, error) {
	query := "SELECT * FROM public.events WHERE id=$1"
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Scan event
	var e Event
	err = stmt.QueryRow(id).Scan(
		&e.ID,
		&e.Name,
		&e.Description,
		&e.Location,
		&e.DateTime,
		&e.UserID,
	)

	if err != nil {
		return nil, err
	}
	return &e, nil
}
