package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"ticketing/ticketing/structs"
	"time"
)

func GetByTransactionId(db *sql.DB, transactionId int) (err error, result structs.TransactionGet) {
	sqlQuery := `SELECT t.id, t.qr_code, t.created_at, c.full_name, c.email, c.phone_number, tic."name" , tic ."date", tic.price, e."name" FROM transaction t 
         		INNER JOIN customer c on c.id = t.customer_id
                INNER JOIN ticket tic on tic.id = t.ticket_id
                INNER JOIN "event" e  on tic.event_id  = e.id 
				WHERE t.id = $1`
	var transaction = structs.TransactionGet{}
	rows, err := db.Query(sqlQuery, transactionId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&transaction.ID,
			&transaction.QrCode,
			&transaction.CreatedAt,
			&transaction.CustomerName,
			&transaction.CustomerEmail,
			&transaction.CustomerPhoneNumber,
			&transaction.TicketName,
			&transaction.TicketDate,
			&transaction.Price,
			&transaction.EventName,
		)
		if err != nil {
			panic(err)
		}
		result = transaction
		return nil, transaction
	}
	err = errors.New("transaction with ID : " + strconv.Itoa(transactionId) + " not found")
	return err, transaction
}

func InsertTransaction(db *sql.DB, transaction structs.Transaction) (structs.Transaction, []error) {
	var errs []error
	sqlQuery := `INSERT INTO transaction (date, qr_code, created_at, updated_at, ticket_id, customer_id) 
				VALUES ($1, $2, $3, $4, $5, $6) 
				Returning *`
	err := db.QueryRow(sqlQuery,
		transaction.Date,
		transaction.QrCode,
		time.Now(),
		time.Now(),
		transaction.TicketId,
		transaction.CustomerId).Scan(
		&transaction.ID,
		&transaction.Date,
		&transaction.QrCode,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.TicketId,
		&transaction.CustomerId)
	if err != nil {
		errs = append(errs, err)
		return transaction, errs
	}
	return transaction, nil
}
