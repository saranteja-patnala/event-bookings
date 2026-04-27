package db

import (
	"database/sql"
	"event-bookings/models/events"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "event_bookings.db")

	if err != nil {
		panic("Unable to establish database connection: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(2)

	CreateTables()
}

func CreateTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL
	);
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}

}

func GetAllEvents() []events.Event {
	query := "select * from events"
	rows, err := DB.Query(query)

	if err != nil {
		panic("Failed to query events: " + err.Error())
	}

	defer rows.Close()
	var eventsArray = []events.Event{}
	for rows.Next() {
		eventData := events.Event{}
		err := rows.Scan(&eventData.ID, &eventData.Name, &eventData.Description,
			&eventData.Location, &eventData.DateTime, &eventData.UserId)
		if err != nil {
			panic("Failed to scan event data: " + err.Error())
		}
		eventsArray = append(eventsArray, eventData)
	}
	return eventsArray
}

func GetEventById(id int64) (events.Event, error) {
	query := "select * from events where id = ?"
	row := DB.QueryRow(query, id)

	if row == nil || row.Err() != nil {
		return events.Event{}, row.Err()
	}
	eventData := events.Event{}
	err := row.Scan(&eventData.ID, &eventData.Name, &eventData.Description,
		&eventData.Location, &eventData.DateTime, &eventData.UserId)
	if err != nil {
		return events.Event{}, err
	}
	return eventData, nil
}

func SaveEvent(event events.Event) events.Event {
	insertQuery := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := DB.Exec(insertQuery, event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		panic("Failed to insert event: " + err.Error())
	}

	id, err := result.LastInsertId()
	event.ID = int(id)
	return event
}

func Save(event events.Event) events.Event {
	insertQuery := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := DB.Exec(insertQuery, event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		panic("Failed to insert event: " + err.Error())
	}

	id, err := result.LastInsertId()
	event.ID = int(id)
	return event
}
