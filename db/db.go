package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initdb() *sql.DB {
	connStr := "user=postgres password=postgres dbname=event-booking sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables(DB)
	return DB
}

func createTables(db *sql.DB) {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS public.events(
		id SERIAL NOT NULL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INT
	)
	`
	_, err := db.Exec(createEventTable)
	if err != nil {
		panic("Could not create events !")
	}
}

func PrepareDB(db *sql.DB, query string) (*sql.Stmt, error) {
	return db.Prepare(query)
}

func RunQuery(db *sql.DB, sql string) (*sql.Rows, error) {
	rows, err := db.Query(sql)
	if err != nil {
		return rows, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}
