package structs

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at""`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryId  int       `json:"category_id"`
}

type EventRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	CategoryId  int    `json:"category_id"`
}
