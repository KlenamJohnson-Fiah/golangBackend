package db

import (
	"bank_transaction/src/domain/model"
	"bank_transaction/src/domain/repository"
	"context"
	"database/sql"
	"log"
	"net/http"
)

type accountRepository struct {
	db *sql.DB
}

func NewMySqlAccountRepository(db *sql.DB) repository.TransactionRepository {
	return &accountRepository{
		db: db,
	}
}
func (d *accountRepository) CheckAccountNumber(ctx context.Context, accNumber string) (bool, error) {
	stmt, err := d.db.Prepare(queryToCheckAccountBalance)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var accountDetails model.AccountBalance

	if checkErr := stmt.QueryRow(accNumber).Scan(&accountDetails.AccountNumber, &accountDetails.CurrentBalance); checkErr != nil {
		return false, checkErr
	}
	return true, nil
}

func (d *accountRepository) Deposit(ctx context.Context, accNumber string, deposit float64) (bool, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.PrepareContext(ctx, updateDeposit)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(deposit, accNumber)
	if updateErr != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return true, nil

}

func (d *accountRepository) Withdraw(ctx context.Context, accNumber string, deposit float64) (bool, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.PrepareContext(ctx, updateDeposit)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(deposit, accNumber)
	if updateErr != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	return true, nil

}

func (d *accountRepository) Transfer(ctx context.Context) {
	panic(http.StatusNotImplemented)
}

func (d *accountRepository) CheckBalance(ctx context.Context, accNumber string) (*model.AccountBalance, error) {
	stmt, err := d.db.Prepare(queryToCheckAccountBalance)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var accountBalance model.AccountBalance

	balanceCheckErr := stmt.QueryRow(accNumber).Scan(&accountBalance.AccountNumber, &accountBalance.CurrentBalance)
	if balanceCheckErr != nil {
		return nil, balanceCheckErr
	}

	return &accountBalance, nil

}
