package dto

import "time"

type TicketGet struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Quota     int       `json:"quota"`
	Price     string    `json:"price"`
	EventId   int       `json:"event_id"`
	EventName string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at""`
	UpdatedAt time.Time `json:"updated_at"`
}
