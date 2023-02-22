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
