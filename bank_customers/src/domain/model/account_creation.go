package model

import "time"

type Account struct {
	AccountNumber   string    `json:"account_number"`
	OwnerID         string    `json:"owner_id"`
	AccountType     string    `json:"account_type"`
	AllowedCurrency string    `json:"allowed_currency"`
	CurrentBalance  float64   `json:"current_balance"`
	DateCreated     time.Time `json:"date_created"`
}
