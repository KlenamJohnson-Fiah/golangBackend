package repository

import (
	"bank_employee/src/domain/model"
)

type EmployeeRepository interface {
	Create(employee *model.Employee) error
	GetByEmail(email string) (*model.Employee, error)
	GetByFirstName(firstname string) (*model.Employee, error)
	Login(email string) (*model.EmployeeLoginSuccessful, error)
}
