package structs

import "time"

type Customer struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	BirthDate   time.Time `json:"birth_date"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsAdmin     bool      `json:"is_admin"`
	Token       string
}

type CustomerRequest struct {
	FullName    string `json:"full_name"`
	BirthDate   string `json:"birth_date"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
}

type CustLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
