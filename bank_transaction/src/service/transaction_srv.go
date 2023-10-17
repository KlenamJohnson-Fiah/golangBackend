package service

import (
	"bank_transaction/src/domain/model"
	"bank_transaction/src/domain/repository"
	"bank_transaction/src/utils/errors"
	"context"
	"fmt"
	"strings"
)

type TransactionServiceInterface interface {
	CheckBalanceService(context.Context, string) (*model.AccountBalance, *errors.Error)
	Deposit(context.Context, string, float64) (*model.AccountBalance, *errors.Error)
	Withdraw(context.Context, string, float64) (*model.AccountBalance, *errors.Error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionServiceInterface {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (t *transactionService) CheckBalanceService(ctx context.Context, accNumber string) (*model.AccountBalance, *errors.Error) {
	accountNumber := strings.TrimSpace(accNumber)

	accountBalance, err := t.transactionRepo.CheckBalance(ctx, accountNumber)
	if err != nil {
		return nil, errors.ErrorNotFound(fmt.Sprintln(err))
	}

	return accountBalance, nil
}

func (t *transactionService) Deposit(ctx context.Context, accNumber string, amount float64) (*model.AccountBalance, *errors.Error) {
	verify, err := t.transactionRepo.CheckAccountNumber(ctx, strings.TrimSpace(accNumber))
	if err != nil && !verify {
		return nil, errors.ErrorBadRequest(fmt.Sprint(err))
	}
	//var getCurrentBalance model.AccountBalance
	getCurrentBalance, checkErr := t.CheckBalanceService(ctx, accNumber)
	if err != nil {
		return nil, checkErr

	}

	newBalance := deposit(getCurrentBalance.CurrentBalance, amount)

	_, depositErr := t.transactionRepo.Deposit(ctx, accNumber, newBalance)
	if depositErr != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(err))
	}

	checkBalanceAfterDeposit, checkErr := t.CheckBalanceService(ctx, accNumber)
	if err != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(checkErr))
	}

	return checkBalanceAfterDeposit, nil
}

func (t *transactionService) Withdraw(ctx context.Context, accNumber string, amount float64) (*model.AccountBalance, *errors.Error) {
	_, err := t.transactionRepo.CheckAccountNumber(ctx, strings.TrimSpace(accNumber))
	if err != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(err))
	}
	//var getCurrentBalance model.AccountBalance
	getCurrentBalance, checkErr := t.CheckBalanceService(ctx, accNumber)
	if err != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(checkErr))

	}

	newBalance := withdraw(getCurrentBalance.CurrentBalance, amount)

	_, withdrawalErr := t.transactionRepo.Withdraw(ctx, accNumber, newBalance)
	if withdrawalErr != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(err))
	}

	checkBalanceAfterWithdrawal, checkErr := t.CheckBalanceService(ctx, accNumber)
	if err != nil {
		return nil, errors.ErrorBadRequest(fmt.Sprint(checkErr))
	}

	return checkBalanceAfterWithdrawal, nil
}
