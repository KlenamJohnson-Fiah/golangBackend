package db

import (
	"bank_customers/src/domain/model"
	"bank_customers/src/domain/repository"
	"bank_customers/src/utils/errors"
	"bank_customers/src/utils/logger"
	"database/sql"
)

type customerRepository struct {
	db *sql.DB
}

func NewMySqlCustomerRepository(db *sql.DB) repository.CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) Create(customer *model.Customers) *errors.RestErr {
	stmt, createErr := c.db.Prepare(queryInsertUser)
	if createErr != nil {
		logger.Error("Error when trying to prepare SAVE user statement", createErr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(customer.Id, customer.FirstName, customer.LastName, customer.Email, customer.Address,
		customer.Status, customer.Password, customer.DateCreated)
	if execErr != nil {
		logger.Error("Error when trying to CREATE user", execErr)
		return errors.NewInternalServerError("Error when trying to create user")
	}

	return nil
}

func (c *customerRepository) GetByEmail(email string) (*model.CustomerDetailsOutput, *errors.RestErr) {
	stmt, getErr := c.db.Prepare(queryGetUserByEmail)
	if getErr != nil {
		logger.Error("Error when trying to prepare GET user statement", getErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	var customer model.CustomerDetailsOutput

	execErr := c.db.QueryRow(queryGetUserByEmail, email).Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Address, &customer.Status)
	if execErr != nil {
		logger.Error("Error when trying to GET user", execErr)
		return nil, errors.NewInternalServerError("Error when trying to GET user by email")
	}

	return &customer, nil
}

func (c *customerRepository) GetByFirstName(firstname string) (*model.MultipleCustomersDetailsOutput, *errors.RestErr) {
	stmt, getErr := c.db.Prepare(queryGetUserByFirstName)
	if getErr != nil {
		logger.Error("Error when trying to prepare GET user statement", getErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	//customer := new(model.Customers)
	row, err := stmt.Query(firstname)
	if err != nil {
		return nil, errors.NewInternalServerError("Error with query")
	}
	matchingCustomersWithFirstName := make(model.MultipleCustomersDetailsOutput, 0)
	for row.Next() {
		var customer model.CustomerDetailsOutput
		execErr := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Address, &customer.Status)
		if execErr != nil {
			logger.Error("Error when trying to GET user", execErr)
			return nil, errors.NewInternalServerError("Error when trying to GET user by firstname")
		}

		matchingCustomersWithFirstName = append(matchingCustomersWithFirstName, customer)

	}
	row.Close()

	// if len(matchingCustomersWithFirstName) == 0 {
	// 	return nil, errors.NewNotFoundError("No customers with the above user name")

	// }

	return &matchingCustomersWithFirstName, nil

}

func (c *customerRepository) Update(customer *model.CustomerDetailsOutput) *errors.RestErr {
	stmt, updateErr := c.db.Prepare(queryUpdateCustomer)
	if updateErr != nil {
		logger.Error("Error when trying to prepare UPDATE user statement", updateErr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err := stmt.Exec(customer.FirstName, customer.LastName, customer.Email, customer.Address, customer.Email)
	if err != nil {
		logger.Error("Error when trying to update customer", err)
		return errors.NewInternalServerError("update error")
	}

	return nil
}

func (c *customerRepository) Login(email string) (*model.CustomerLoginSuccessful, *errors.RestErr) {
	stmt, getErr := c.db.Prepare(queryLoginByEmail)
	if getErr != nil {
		logger.Error("Error when trying to prepare GET user statement", getErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	var customer model.CustomerLoginSuccessful
	execErr := stmt.QueryRow(email).Scan(&customer.Id, &customer.Email, &customer.Password)
	if execErr != nil {
		logger.Error("Error when trying to GET user", execErr)
		return nil, errors.NewInternalServerError("Error when trying to GET user by email")
	}

	return &customer, nil
}

//Secondary DB connection for creating an account when user is created
func (c *customerRepository) CreateAccount(account *model.Account) *errors.RestErr {
	stmt, createErr := c.db.Prepare(queryCreateUserAccount)
	if createErr != nil {
		logger.Error("Error when trying to prepare CreatAccount statement", createErr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, execErr := stmt.Exec(account.AccountNumber, account.OwnerID, account.AccountType, account.AllowedCurrency, account.CurrentBalance, account.DateCreated)
	if execErr != nil {
		logger.Error("Error when trying to CREATE account", execErr)
		return errors.NewInternalServerError("Error when trying to create account")
	}

	return nil

}
