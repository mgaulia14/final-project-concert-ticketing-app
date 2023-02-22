package repository

import (
	"database/sql"
	"errors"
	"final-project-ticketing-api/structs"
	"strconv"
	"time"
)

func GetByTicketId(db *sql.DB, ticketId int) (err error, result structs.Ticket) {
	sqlQuery := `SELECT * FROM ticket
				WHERE ticket.id = $1`
	var ticket = structs.Ticket{}
	rows, err := db.Query(sqlQuery, ticketId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&ticket.ID,
			&ticket.Name,
			&ticket.Date,
			&ticket.Quota,
			&ticket.Price,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
			&ticket.EventId,
		)
		if err != nil {
			panic(err)
		}
		result = ticket
		return nil, ticket
	}
	err = errors.New("ticket with ID : " + strconv.Itoa(ticketId) + " not found")
	return err, ticket
}

func InsertTicket(db *sql.DB, ticket structs.Ticket) (structs.Ticket, []error) {
	var errs []error
	sqlQuery := `INSERT INTO ticket (name, date, quota, price, event_id, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		ticket.Name,
		ticket.Date,
		ticket.Quota,
		ticket.Price,
		ticket.EventId,
		time.Now(),
		time.Now()).Scan(
		&ticket.ID,
		&ticket.Name,
		&ticket.Date,
		&ticket.Quota,
		&ticket.Price,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
		&ticket.EventId)
	if err != nil {
		errs = append(errs, err)
		return ticket, errs
	}
	return ticket, nil
}

func UpdateTicket(db *sql.DB, ticket structs.Ticket) (structs.Ticket, []error) {
	var errs []error
	sqlQuery := `UPDATE ticket 
				SET name = $1, 
				    date = $2,
				    quota = $3, 
				    price = $4,
				    event_id = $5,
				    updated_at = $6
				WHERE id = $7`

	err := db.QueryRow(sqlQuery,
		ticket.Name,
		ticket.Date,
		ticket.Quota,
		ticket.Price,
		ticket.EventId,
		time.Now(),
		ticket.ID).Scan(
		&ticket.ID,
		&ticket.Name,
		&ticket.Date,
		&ticket.Quota,
		&ticket.Price,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
		&ticket.EventId)
	if errs != nil {
		errs = append(errs, err)
		return ticket, errs
	}
	return ticket, nil
}

func DeleteTicket(db *sql.DB, ticketId int) (err error) {
	sqlQuery := "DELETE FROM ticket WHERE id = $1"
	errs := db.QueryRow(sqlQuery, ticketId)
	return errs.Err()
}
