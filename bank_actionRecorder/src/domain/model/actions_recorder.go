package model

import "time"

type Recorder struct {
	//Action_ID  int64     `json:"action_id"`
	CustomerID string    `json:"customer_id"`
	Action     string    `json:"action"`
	Amount     float64   `json:"amount"`
	ActionDate time.Time `json:"action_date"`
}

type CustomerActivityActionRecorder struct {
	ActivityID       string    `json:"activity_id"`
	CustomerID       string    `json:"customer_id"`
	ActivityDateTime time.Time `json:"activity_date_time"`
	ProfileAction    string    `json:"profile_action"`
}

type EmployeeActivityActionRecorder struct {
	ActivityID       string    `json:"activity_id"`
	EmployeeID       string    `json:"employee_id"`
	ActivityDateTime time.Time `json:"activity_date_time"`
	ActionPerformed  string    `json:"action_performed"`
	Amount           float64   `json:"amount"`
	ProfileAction    string    `json:"profile_action"`
}

type TransactionActivityActionRecorder struct {
	TransactionID    string    `json:"transaction_id"`
	CustomerID       string    `json:"customer_id"`
	AmountBeforeT    float64   `json:"amount_before_T"`
	AmountTransfered float64   `json:"amount_transferred"`
	AmountAfterT     float64   `json:"amount_after_T"`
	DateAndTime      time.Time `json:"date_and_time"`
}
