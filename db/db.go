package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initdb() {
	connStr := "user=agt dbname=event-booking sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables(DB)
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

// func RunQuery(sql string) (*sql.Rows, error) {
// 	rows, err := DB.Query(sql)
// 	if err != nil {
// 		log.Fatalf("Failed to execute query: %v", err)
// 	}
// 	return rows, err
// }

// func FetchRows(rows *sql.Rows) []any {
// 	result := []any{}
// 	for rows.Next() {
// 		var user_id int
// 		var username, email string

// 		// Scan the columns into variables
// 		err := rows.Scan(&user_id, &username, &email)
// 		if err != nil {
// 			log.Fatalf("Failed to scan row: %v", err)
// 		}
// 		row := make(map[string]any, 3)
// 		row["id"] = user_id
// 		row["username"] = username
// 		row["email"] = email
// 		result = append(result, row)
// 	}
// 	return result
// }
