package http

import (
	"bank_transaction/src/domain/model"
	"bank_transaction/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
	CheckBalanceHandler(c *gin.Context)
	DepositeHandler(c *gin.Context)
	WithdrawHandler(c *gin.Context)
}

type transactionHandler struct {
	transactionServices service.TransactionServiceInterface
}

func NewTransactionHandler(transactionServices service.TransactionServiceInterface) TransactionHandler {
	return &transactionHandler{
		transactionServices: transactionServices,
	}
}

func (t *transactionHandler) CheckBalanceHandler(c *gin.Context) {
	accountNumber := c.Param("accountNumber")

	accountBalance, err := t.transactionServices.CheckBalanceService(c, accountNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return

	}
	c.JSON(http.StatusOK, accountBalance)

}

func (t *transactionHandler) DepositeHandler(c *gin.Context) {
	var accountDeposit model.AccountDeposit

	if err := c.ShouldBindJSON(&accountDeposit); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	checkBalanceAfterDeposit, err := t.transactionServices.Deposit(c, accountDeposit.AccountNumber, accountDeposit.DepositAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, checkBalanceAfterDeposit)

}

func (t *transactionHandler) WithdrawHandler(c *gin.Context) {

	var accountWithdrawal model.AccountDeposit

	if err := c.ShouldBindJSON(&accountWithdrawal); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	checkBalanceAfterWithdrawal, err := t.transactionServices.Withdraw(c, accountWithdrawal.AccountNumber, accountWithdrawal.DepositAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, checkBalanceAfterWithdrawal)

}
