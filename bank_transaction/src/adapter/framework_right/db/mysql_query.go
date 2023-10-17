package db

const (
	queryToCheckAccountNumber  = "SELECT account_number,owner_id,account_type,allowed_currency,current_balance FROM customer_bank_accounts WHERE account_number=?;"
	queryToCheckAccountBalance = "SELECT account_number,current_balance FROM customer_bank_accounts WHERE account_number=?;"

	updateDeposit = "UPDATE customer_bank_accounts SET current_balance=? WHERE account_number=?;"
)
