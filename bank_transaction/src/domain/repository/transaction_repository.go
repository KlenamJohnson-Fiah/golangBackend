package repository

import (
	"bank_transaction/src/domain/model"
	"context"
)

type TransactionRepository interface {
	CheckAccountNumber(context.Context, string) (bool, error)
	Deposit(context.Context, string, float64) (bool, error)
	Withdraw(context.Context, string, float64) (bool, error)
	Transfer(context.Context)
	CheckBalance(context.Context, string) (*model.AccountBalance, error)
}
