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

type Transfer struct {
	AccountNumber         string    `json:"account_number"`
	AllowedCurrency       string    `json:"allowed_currency"`
	CurrentBalance        float64   `json:"current_balance"`
	ReceiverAccountNumber int64     `json:"reciever_account_number"`
	TransferDate          time.Time `json:"transfer_date"`
}

type AccountBalance struct {
	AccountNumber  string  `json:"account_number"`
	CurrentBalance float64 `json:"current_balance"`
}

type AccountDeposit struct {
	AccountNumber string  `json:"account_number"`
	DepositAmount float64 `json:"current_balance"`
}
