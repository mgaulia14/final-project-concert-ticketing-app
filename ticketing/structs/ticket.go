package structs

import "time"

type Ticket struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Quota     int       `json:"quota"`
	Price     string    `json:"price"`
	EventId   int       `json:"event_id"`
	CreatedAt time.Time `json:"created_at""`
	UpdatedAt time.Time `json:"updated_at"`
}

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

type TicketRequest struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Quota   int    `json:"quota"`
	Price   string `json:"price"`
	EventId int    `json:"event_id"`
}
