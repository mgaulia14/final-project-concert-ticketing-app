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
	WalletId    int       `json:"wallet_id"`
}

type CustomerRequest struct {
	FullName    string `json:"full_name"`
	BirthDate   string `json:"birth_date"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
