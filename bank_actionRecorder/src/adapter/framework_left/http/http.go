package http

import (
	"bank_actionRecorder/src/service"
	"fmt"
	"net/http"

	"bank_actionRecorder/src/domain/model"
	"bank_actionRecorder/src/utils/errors"

	"github.com/gin-gonic/gin"
)

type ActionsRecorderHandler interface {
	AddCustomerActionHandler(c *gin.Context)
	RetrieveCustomerActionHandler(c *gin.Context)
	AddEmployeeActionHandler(c *gin.Context)
	RetrieveEmployeeActionHandler(c *gin.Context)
	AddTransactionActivityRecord(c *gin.Context)
	RetrieveTransactionActivityRecord(c *gin.Context)
}

type actionRecorderHandler struct {
	actionRecorderService service.ActionRecorderServiceInterface
}

func NewActionRecorderHandler(actionRecorderService service.ActionRecorderServiceInterface) ActionsRecorderHandler {
	return &actionRecorderHandler{
		actionRecorderService: actionRecorderService,
	}
}

func (a *actionRecorderHandler) AddCustomerActionHandler(c *gin.Context) {
	var customerActivity model.CustomerActivityActionRecorder

	if err := c.ShouldBindJSON(&customerActivity); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Printf("%+v\n", customerActivity)
		return

	}
	if err := a.actionRecorderService.AddCustomerActivityRecord(c, &customerActivity); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, customerActivity)
	return

}

func (a *actionRecorderHandler) RetrieveCustomerActionHandler(c *gin.Context) {
	id := c.Param("id")

	customerActivity, err := a.actionRecorderService.RetrieveCustomerActivityRecord(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

	}
	c.JSON(http.StatusOK, customerActivity)

}

func (a *actionRecorderHandler) AddEmployeeActionHandler(c *gin.Context) {
	var employeeActivity model.EmployeeActivityActionRecorder

	if err := c.ShouldBindJSON(&employeeActivity); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		fmt.Printf("%+v\n", employeeActivity)
		return

	}
	if err := a.actionRecorderService.AddEmployeeActivityRecord(c, &employeeActivity); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, employeeActivity)
	return

}

func (a *actionRecorderHandler) RetrieveEmployeeActionHandler(c *gin.Context) {
	id := c.Param("id")

	employeeActivity, err := a.actionRecorderService.RetrieveEmployeeActivityRecord(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)

	}
	c.JSON(http.StatusOK, employeeActivity)

}

func (a *actionRecorderHandler) AddTransactionActivityRecord(c *gin.Context) {

	//TODO: Get the handler for Transactions

}

func (a *actionRecorderHandler) RetrieveTransactionActivityRecord(c *gin.Context) {

	//TODO: Get the handler for Transactions
}
