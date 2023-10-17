package http

import (
	"bank_customers/src/domain/model"
	"bank_customers/src/service"
	"fmt"
	"net/http"

	"bank_customers/src/utils/errors"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	Create(c *gin.Context)
	GetCustomerByEmail(c *gin.Context)
	GetCustomerByFirstName(c *gin.Context)
	UpdateCustomerInfo(c *gin.Context)
	LoginByEmail(c *gin.Context)
}

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) CustomerHandler {
	return &customerHandler{
		customerService: customerService,
	}
}

func (customer *customerHandler) Create(c *gin.Context) {
	var bankCustomer model.Customers

	if err := c.ShouldBindJSON(&bankCustomer); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	customer.customerService.CreateCustomer(c, &bankCustomer)

	c.JSON(http.StatusCreated, bankCustomer)

}

func (customer *customerHandler) GetCustomerByEmail(c *gin.Context) {
	customerEmail, emailErr := customer.customerService.GetCustomerByEmail(c, c.Param("email"))
	if emailErr != nil {
		c.JSON(emailErr.Status, emailErr)
	}

	c.JSON(http.StatusFound, customerEmail)

}

func (customer *customerHandler) GetCustomerByFirstName(c *gin.Context) {
	customerFirstname, firstnameErr := customer.customerService.GetCustomerByFirstName(c, c.Param("firstname"))
	if firstnameErr != nil {
		c.JSON(firstnameErr.Status, firstnameErr)
	}

	c.JSON(http.StatusFound, customerFirstname)

}

func (customer *customerHandler) UpdateCustomerInfo(c *gin.Context) {
	var customerUpdate model.CustomerDetailsOutput
	_, getErr := customer.customerService.GetCustomerByEmail(c, c.Param("email"))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)

	}

	if err := c.ShouldBindJSON(&customerUpdate); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Printf("%+v\n", customerUpdate)
	}
	customer.customerService.UpdateCustomer(c, &customerUpdate)

	c.JSON(http.StatusCreated, customerUpdate)

}

func (customer *customerHandler) LoginByEmail(c *gin.Context) {
	//TODO: Fix issue with binding of data

	var customerCredentails model.CustomerLoginCredentials
	if err := c.ShouldBind(&customerCredentails); err != nil {
		loginErr := errors.NewInternalServerError("invalid json body")
		c.JSON(loginErr.Status, loginErr)

	}
	fmt.Println(customerCredentails)
	resp, err := customer.customerService.LoginCustomer(c, customerCredentails.Email, customerCredentails.Password)
	if err != nil {
		c.JSON(err.Status, err)
		return

	}

	c.JSON(http.StatusOK, resp)

}
