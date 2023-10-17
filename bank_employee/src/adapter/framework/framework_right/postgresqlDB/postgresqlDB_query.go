package postgresqldb

var (
	createEmployeeDB    = "CREATE DATABASE employeeDB;"
	selectEmployeeDB    = "USE employeeDB;"
	createEmployeeTable = "CREATE TABLE employees IF NOT EXIST(id VARCHAR (60) PRIMARY KEY,first_name VARCHAR (30) NOT NULL,last_name VARCHAR (30) NOT NULL,email VARCHAR (30) UNIQUE NOT NULL,password VARCHAR (30) NOT NULL,role VARCHAR (10),branch VARCHAR (30));"

	queryToCreateEmployee = "INSERT INTO employees (id,first_name,last_name,email,password,role,branch,status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8);"
	queryByEmail          = "SELECT id,email,password FROM employees WHERE email=$1;"
	queryByFirstName      = "SELECT id,first_name,last_name,email,role,branch FROM employees WHERE first_name=$1;"
)
