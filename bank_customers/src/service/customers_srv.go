package service

import (
	"bank_customers/src/domain/model"
	"bank_customers/src/domain/repository"
	accountnumbergenerator "bank_customers/src/utils/accountNumber_generator"
	"bank_customers/src/utils/errors"
	"bank_customers/src/utils/httpclient"
	"bank_customers/src/utils/logger"
	"bank_customers/src/utils/password_hash"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

const (
	accountCreation      = "accountCreation"
	accountDetailsUpdate = "accountDetailsUpdate"
)

type CustomerService interface {
	CreateCustomer(context.Context, *model.Customers) *errors.RestErr
	GetCustomerByEmail(context.Context, string) (*model.CustomerDetailsOutput, *errors.RestErr)
	GetCustomerByFirstName(context.Context, string) (*model.MultipleCustomersDetailsOutput, *errors.RestErr)
	UpdateCustomer(context.Context, *model.CustomerDetailsOutput) *errors.RestErr
	LoginCustomer(context.Context, string, string) (*model.CustomerLoginSuccessful, *errors.RestErr)
}

type customerService struct {
	customerRepo repository.CustomerRepository
	//accountCreation db.AccountCreationInterface
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (c *customerService) CreateCustomer(ctx context.Context, customer *model.Customers) *errors.RestErr {
	if ctx.Value("role") != "employee" {
		return errors.NewBadRequestError("user role unauthorized")
	}
	customer.Id = uuid.New().String()
	customer.DateCreated = time.Now().UTC()
	customer.Status = "active"
	customer.Password = password_hash.HashAndSalt([]byte(customer.Password))
	if !password_hash.EmailValid(customer.Email) {
		return errors.NewBadRequestError("email structure invalid")
	}

	fmt.Printf("%T\n"+"%+v\n"+"%v\n"+"%v\n", customer, customer, customer == nil, c)
	if err := c.customerRepo.Create(customer); err != nil {
		return errors.NewInternalServerError("error create")
	}

	_, err := httpclient.HTTPClientPost("http://localhost:8081/actions/customer/record", map[string]interface{}{
		"activity_id":        shortuuid.New(),
		"customer_id":        customer.Id,
		"activity_date_time": customer.DateCreated,
		"profile_action":     accountCreation,
	})

	if err != nil {
		errors.NewInternalServerError("error tracking transaction")

		return nil
	}

	accountCreationErr := c.customerRepo.CreateAccount(
		&model.Account{
			AccountNumber:   accountnumbergenerator.RandomString(8),
			OwnerID:         customer.Id,
			AccountType:     "saving",
			AllowedCurrency: "GH",
			CurrentBalance:  0.00,
			DateCreated:     time.Now().UTC(),
		},
	)
	if accountCreationErr != nil {
		return accountCreationErr
	}

	return nil
}

func (c *customerService) GetCustomerByEmail(ctx context.Context, email string) (*model.CustomerDetailsOutput, *errors.RestErr) {
	result, err := c.customerRepo.GetByEmail(email)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil

}

func (c *customerService) GetCustomerByFirstName(ctx context.Context, firstname string) (*model.MultipleCustomersDetailsOutput, *errors.RestErr) {
	result, err := c.customerRepo.GetByFirstName(firstname)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil

}

//TODO: Crashes if changes are made to email.
//TODO: Proper JSON Presentation of data
func (c *customerService) UpdateCustomer(ctx context.Context, customer *model.CustomerDetailsOutput) *errors.RestErr {
	customerFinder, err := c.GetCustomerByEmail(ctx, customer.Email)
	if err != nil {
		logger.Info("no user with this email to update")
		return err
	} else {
		customerFinder.Id = customer.Id
		customerFinder.FirstName = customer.FirstName
		customerFinder.LastName = customer.LastName
		customerFinder.Address = customer.Address
		customerFinder.Email = customer.Email
		customerFinder.DateLastModified = time.Now().UTC()
	}

	if updateErr := c.customerRepo.Update(customerFinder); updateErr != nil {
		return updateErr
	}
	_, updateErr := httpclient.HTTPClientPost("http://localhost:8081/actions/customer/record", map[string]interface{}{
		"activity_id":        shortuuid.New(),
		"customer_id":        customerFinder.Id,
		"activity_date_time": time.Now().UTC(),
		"profile_action":     accountDetailsUpdate,
	})
	if updateErr != nil {
		return updateErr
	}
	return nil

}

func (c *customerService) LoginCustomer(ctx context.Context, email, password string) (*model.CustomerLoginSuccessful, *errors.RestErr) {
	customerCredentials, err := c.customerRepo.Login(email)
	if err != nil {
		logger.Info("Can't login. possible error in credentials")
		return nil, err
	}
	if !password_hash.ComparePasswords(customerCredentials.Password, []byte(password)) {
		return nil, errors.NewBadRequestError("Password doesn't match")
	}
	fmt.Println(customerCredentials)
	_, loginErr := httpclient.HTTPClientPost("http://localhost:8081/actions/customer/record", map[string]interface{}{
		"activity_id":        shortuuid.New(),
		"customer_id":        customerCredentials.Id,
		"activity_date_time": time.Now().UTC(),
		"profile_action":     accountDetailsUpdate,
	})
	if loginErr != nil {
		return nil, loginErr
	}

	return customerCredentials, nil
}
