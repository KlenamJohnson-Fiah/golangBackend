package db

const (
	//INSERT TO DBs
	insertCustomerActivityRecorder = "INSERT INTO customer_activity_recorder(activity_id,customer_id,activity_date_time,profile_action) VALUES(?, ?, ?, ?);"
	insertEmployeeActivityRecorder = "INSERT INTO employee_activity_recorder(activity_id,employee_id,activity_date_time,action_performed,amount,profile_action) VALUES(?, ?, ?, ?,?,?);"
	insertTransactionRecorder      = "INSERT INTO transaction_activity_recorder(transaction_id,customer_id,amount_before_T,amount_transferred,amount_after_T,date_and_time) VALUES(?, ?, ?, ?,?);"

	//QUERY DBs
	queryCustomerActivityRecorder = "SELECT (activity_id,customer_id,activity_date_time,profile_action)FROM employee_activity_recorder WHERE activity_id=?;"
	queryEmployeeActivityRecorder = "SELECT (activity_id,employee_id,activity_date_time,action_performed,amount,profile_action) FROM employee_activity_recorder WHERE activity_id=?;"
	queryTransactionRecorder      = "SELECT (transaction_id,amount_before_T,amount_transferred,amount_after_T,date_and_time) FROM transaction_activity_recorder WHERE transaction_id=?;"
)
