package repository

import (
	"database/sql"
	"errors"
	"final-project-ticketing-api/structs"
	"strconv"
	"time"
)

func GetAllEvent(db *sql.DB) (err error, results []structs.EventGet) {
	sqlQuery := `SELECT e.id, e.name, e.description, e.start_date, e.end_date, e.category_id, c.name 
				FROM event e 
				INNER JOIN category c on c.id = e.category_id`
	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var event = structs.EventGet{}
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.CategoryId,
			&event.CategoryName,
		)
		if err != nil {
			panic(err)
		}
		results = append(results, event)
	}
	return
}

func GetEventById(db *sql.DB, id int) (err error, result structs.EventGet) {
	sqlQuery := `SELECT e.id, e.name, e.description, e.start_date, e.end_date, e.category_id, c.name 
				FROM event e 
				INNER JOIN category c on c.id = e.category_id
				WHERE e.id = $1`
	var event = structs.EventGet{}
	rows, _ := db.Query(sqlQuery, id)

	if !rows.Next() {
		err = errors.New("event with ID : " + strconv.Itoa(id) + " not found")
		return err, event
	} else {
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.CategoryId,
			&event.CategoryName,
		)
		if err != nil {
			panic(err)
		}
		result = event
	}
	return nil, result
}

func InsertEvent(db *sql.DB, event structs.Event) (structs.Event, error) {
	sqlQuery := `INSERT INTO event (name, description, start_date, end_date, created_at, updated_at, category_id) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		event.Name,
		event.Description,
		event.StartDate,
		event.EndDate,
		time.Now(),
		time.Now(),
		event.CategoryId).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.StartDate,
		&event.EndDate,
		&event.CreatedAt,
		&event.UpdatedAt,
		&event.CategoryId)
	if err != nil {
		return event, err
	}
	return event, nil
}

func UpdateEvent(db *sql.DB, event structs.Event) (structs.Event, []error) {
	var errs []error
	sqlQuery := `UPDATE event 
				SET name = $1,
				    description = $2,
				    start_date = $3,
				    end_date = $4,
				    category_id = $5,
                    updated_at = $6
                WHERE id = $7`
	err := db.QueryRow(sqlQuery,
		event.Name,
		event.Description,
		event.StartDate,
		event.EndDate,
		event.CategoryId,
		time.Now(),
		event.ID).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.StartDate,
		&event.EndDate,
		&event.CategoryId,
		&event.CreatedAt,
		&event.UpdatedAt)
	if errs != nil {
		errs = append(errs, err)
		return event, errs
	}
	return event, nil
}

func DeleteEvent(db *sql.DB, eventId int) (err error) {
	sqlQuery := `DELETE FROM event WHERE id = $1`
	errs := db.QueryRow(sqlQuery, eventId)
	return errs.Err()
}

func GetAllTicketByEventId(db *sql.DB, id int) (err error, results []structs.TicketGet) {
	sqlQuery := `SELECT ticket.id, ticket.name, ticket.date, ticket.quota, ticket.price, ticket.created_at, ticket.updated_at, event.id, event.name
				FROM ticket
				INNER JOIN event
				ON ticket.event_id= event.id
				WHERE ticket.event_id = $1`
	rows, err := db.Query(sqlQuery, &id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ticket = structs.TicketGet{}
		err = rows.Scan(
			&ticket.ID,
			&ticket.Name,
			&ticket.Date,
			&ticket.Quota,
			&ticket.Price,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
			&ticket.EventId,
			&ticket.EventName,
		)
		if err != nil {
			panic(err)
		}
		results = append(results, ticket)
	}
	return
}
