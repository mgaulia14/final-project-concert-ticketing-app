package structs

import "time"

type Wallet struct {
	ID            int       `json:"id"`
	Balance       float64   `json:"balance"`
	AccountName   string    `json:"account_name"`
	AccountNumber int       `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type WalletTopUp struct {
	Balance       float64 `json:"balance"`
	AccountNumber int     `json:"account_number"`
}
