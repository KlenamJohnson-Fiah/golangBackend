package repository

import (
	"bank_customers/src/domain/model"
	"bank_customers/src/utils/errors"
)

type CustomerRepository interface {
	Create(customer *model.Customers) *errors.RestErr
	GetByEmail(email string) (*model.CustomerDetailsOutput, *errors.RestErr)
	GetByFirstName(firstname string) (*model.MultipleCustomersDetailsOutput, *errors.RestErr)
	Update(customer *model.CustomerDetailsOutput) *errors.RestErr
	Login(email string) (*model.CustomerLoginSuccessful, *errors.RestErr)

	CreateAccount(account *model.Account) *errors.RestErr
}
