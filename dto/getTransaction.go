package dto

import "time"

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
