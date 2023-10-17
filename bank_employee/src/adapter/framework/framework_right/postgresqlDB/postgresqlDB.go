package postgresqldb

import (
	"bank_employee/src/domain/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type DBInterface interface {
	Create(pg *sql.DB, employee *model.Employee) error
	GetByEmail(pg *sql.DB, email string) error
	GetByName(pg *sql.DB, name string) (*model.Employee, error)
}

func Create(pg *sql.DB, employee *model.Employee) error {
	stmt, err := pg.Prepare(queryToCreateEmployee)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	_, createEmployeeErr := stmt.Exec(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Password, &employee.Role, &employee.Branch, &employee.Status)
	if createEmployeeErr != nil {
		return createEmployeeErr
	}
	return nil
}

func GetByEmail(pg *sql.DB, credential *model.EmployeeLoginCredentials) (*model.EmployeeLoginSuccessful, error) {
	var employee model.EmployeeLoginSuccessful
	stmt, err := pg.Prepare(queryByEmail)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer stmt.Close()

	Geterr := stmt.QueryRow(credential.Email).Scan(&employee.Id, &employee.Email, &employee.Password)
	if Geterr != nil {
		return nil, sql.ErrNoRows
	}
	return &employee, nil
}

func GetByName(pg *sql.DB, name string) (model.SliceOfEmployees, error) {

	stmt, err := pg.Prepare(queryByFirstName)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	allEmployees := make([]model.Employee, 0)
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Role, &employee.Branch)
		if err != nil {
			log.Println(err)
		}
		allEmployees = append(allEmployees, employee)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return allEmployees, nil
}
