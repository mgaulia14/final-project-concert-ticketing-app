package structs

import (
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	QrCode     string    `json:"qr_code"`
	CreatedAt  time.Time `json:"created_at""`
	UpdatedAt  time.Time `json:"updated_at"`
	TicketId   int       `json:"ticket_id"`
	CustomerId int       `json:"customer_id"`
}

type TransactionRequest struct {
	Date       string `json:"date"`
	TicketId   int    `json:"ticket_id"`
	CustomerId int    `json:"customer_id"`
}

type TransactionGet struct {
	ID                  int       `json:"id"`
	QrCode              string    `json:"qr_code"`
	CreatedAt           time.Time `json:"created_at""`
	CustomerId          string    `json:"customer_id"`
	CustomerName        string    `json:"customer_name"`
	CustomerEmail       string    `json:"email"`
	CustomerPhoneNumber string    `json:"phone_number"`
	TicketId            string    `json:"ticket_id"`
	TicketName          string    `json:"ticket_name"`
	TicketDate          time.Time `json:"ticket_date"`
	Price               string    `json:"price"`
	EventId             string    `json:"event_id"`
	EventName           string    `json:"event_name"`
}
