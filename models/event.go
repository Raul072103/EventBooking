package models

import (
	"EventBooking/db"
	"database/sql"
	"errors"
	"time"
)

type Event struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" binding:"required" db:"name"`
	Description string    `json:"description"  binding:"required" db:"description"`
	Location    string    `json:"location"  binding:"required" db:"location"`
	DateTime    time.Time `json:"date_time"  binding:"required" db:"dateTime"`
	UserID      int64     `json:"user_id" db:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id=?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.DateTime, &event.UserID)

	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var Event Event
		err := rows.Scan(&Event.ID, &Event.Name, &Event.Location, &Event.Description, &Event.DateTime, &Event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, Event)
	}

	return events, nil
}

func Update(event Event) error {
	query := `
	UPDATE events
	SET name=?, description=?, location=?, dateTime=?
	WHERE id=?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			err = errors.Join(err)
		}
	}(stmt)

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			errors.Join(err)
		}
	}(stmt)

	_, err = db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
