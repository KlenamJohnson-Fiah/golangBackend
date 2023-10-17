package service

import (
	postgresqldb "bank_employee/src/adapter/framework/framework_right/postgresqlDB"
	"bank_employee/src/domain/model"
	"database/sql"
	"log"

	"bank_employee/src/utils/password_hash"

	"github.com/twinj/uuid"
)

type EmployeeServiceInterface interface {
	CreateEmployee(employee *model.Employee) (*model.Employee, error)
	Login(pg *sql.DB, credential *model.EmployeeLoginCredentials) (*model.EmployeeLoginSuccessful, error)
	GetByName(pg *sql.DB, name string) (*model.Employee, error)
}

func CreateEmployee(pg *sql.DB, employee *model.Employee) (*model.Employee, error) {
	employee.Id = uuid.NewV4().String()
	employee.Password = password_hash.HashAndSalt([]byte(employee.Password))

	err := postgresqldb.Create(pg, employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func Login(pg *sql.DB, credential *model.EmployeeLoginCredentials) (*model.EmployeeLoginSuccessful, error) {
	employee, err := postgresqldb.GetByEmail(pg, credential)
	if err != nil {
		return nil, err
	}
	passwordValidation := password_hash.ComparePasswords(employee.Password, []byte(credential.Password))
	if passwordValidation != true {
		log.Println("passwords don't match")
		return nil, err
	}
	return employee, nil
}

func GetByName(pg *sql.DB, name string) (*model.SliceOfEmployees, error) {
	employeeName, err := postgresqldb.GetByName(pg, name)
	if err != nil {
		return nil, err
	}

	return &employeeName, nil
}
