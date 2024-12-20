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

func GetAllEvents() ([]Event, error) {
	events = []Event{}

	query := "SELECT * FROM public.events"
	DB := db.Initdb()
	rows, err := db.RunQuery(DB, query)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event

		// Scan the columns into variables
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		events = append(events, e)
	}
	return events, err
}

func GetEvent(id string) (Event, error) {
	var event Event
	query := "SELECT * FROM public.events WHERE id=$1"
	DB := db.Initdb()
	stmt, err := db.PrepareDB(DB, query)
	if err != nil {
		return event, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return event, err
	}
	log.Fatal(result)
	return event, nil
}
