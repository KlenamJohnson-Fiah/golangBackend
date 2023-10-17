package db

const (
	//Customer Query
	queryInsertUser         = "INSERT into customers(id,first_name,last_name,email,address,status,password,date_created) VALUES(?, ?, ?, ?, ?, ?, ?,?);"
	queryGetUserByEmail     = "SELECT id,first_name, last_name, email, address, status FROM customers WHERE email=?;"
	queryGetUserByFirstName = "SELECT id,first_name, last_name, email, address, status FROM customers WHERE first_name=?;"
	queryUpdateCustomer     = "UPDATE customers SET first_name=?,last_name=?,email=?, address=? WHERE email=?;"
	queryLoginByEmail       = "SELECT id,email,password FROM customers WHERE email=?;"
	//queryLoginByEmail       = "SELECT id,first_name, last_name, email, address, status,password FROM customers WHERE email=?;"

	//Action_Recorder Query
	queryInsertActionTaken = "INSERT into actions(customer_id,action,amount,action_date) VALUES(?, ?, ?, ?);"

	//Account Creation
	queryCreateUserAccount = "INSERT into customer_bank_accounts(account_number,owner_id,account_type,allowed_currency,current_balance,date_created) VALUES(?,?,?,?,?,?);"
)
