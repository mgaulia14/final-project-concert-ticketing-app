package dto

import "time"

type EventGet struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
}
