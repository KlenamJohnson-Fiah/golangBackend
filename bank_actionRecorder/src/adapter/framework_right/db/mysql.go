package db

import (
	"bank_actionRecorder/src/domain/model"
	"bank_actionRecorder/src/domain/repository"
	"bank_actionRecorder/src/utils/errors"
	"bank_actionRecorder/src/utils/logger"
	"context"
	"database/sql"
	"fmt"
)

type ActionRecorderRepository struct {
	ActionDB *sql.DB
}

func NewActionRecorderRepository(actionDB *sql.DB) repository.ActionRecorderRepository {
	return &ActionRecorderRepository{
		ActionDB: actionDB,
	}
}

// func (adb *ActionRecorderRepository) AddRecord(action *model.Recorder) *errors.RestErr {
// 	stmt, err := adb.ActionDB.Prepare(queryInsertActionTaken)
// 	if err != nil {
// 		logger.Error("error preparing the AddRecord query", err)
// 		return errors.NewInternalServerError("can't add record")
// 	}
// 	defer stmt.Close()

// 	_, exeErr := stmt.Exec(action.CustomerID, action.Action, action.Amount, action.ActionDate)
// 	if exeErr != nil {
// 		logger.Error("error execting the AddRecord exec", exeErr)
// 		return errors.NewInternalServerError("can't add record")
// 	}
// 	return nil
// }

func (a *ActionRecorderRepository) AddCustomerActivityRecord(ctx context.Context, action *model.CustomerActivityActionRecorder) *errors.RestErr {
	stmt, err := a.ActionDB.Prepare(insertCustomerActivityRecorder)
	if err != nil {
		logger.Error("error preparing the AddCustomerRecord query", err)
		return errors.NewInternalServerError("can not add record")
	}
	defer stmt.Close()

	_, exeErr := stmt.Exec(action.ActivityID, action.CustomerID, action.ActivityDateTime, action.ProfileAction)
	if exeErr != nil {
		logger.Error("error execting the AddRecord exec", exeErr)
		return errors.NewInternalServerError("can't add record")
	}
	return nil
}

func (a *ActionRecorderRepository) RetrieveCustomerActivityRecord(ctx context.Context, id string) (*model.CustomerActivityActionRecorder, *errors.RestErr) {
	stmt, err := a.ActionDB.Prepare(queryCustomerActivityRecorder)
	if err != nil {
		logger.Error("error preparing the QueryCustomerRecord query", err)
		return nil, errors.NewInternalServerError("can not add record")
	}
	defer stmt.Close()

	var customer *model.CustomerActivityActionRecorder
	if exeErr := stmt.QueryRow(id).Scan(customer.ActivityID, customer.CustomerID, customer.ActivityDateTime, customer.ProfileAction); err != nil {
		logger.Error(fmt.Sprintf("error querying for activities of user with ID: %s\n", id), exeErr)
		return nil, errors.NewInternalServerError("can't query customer activity")
	}

	return customer, nil
}
func (a *ActionRecorderRepository) AddEmployeeActivityRecord(ctx context.Context, action *model.EmployeeActivityActionRecorder) *errors.RestErr {
	stmt, err := a.ActionDB.Prepare(insertEmployeeActivityRecorder)
	if err != nil {
		logger.Error("error preparing the AddEmployeeRecord query", err)
		return errors.NewInternalServerError("can not add record")
	}
	defer stmt.Close()

	_, exeErr := stmt.Exec(action.ActivityID, action.EmployeeID, action.ActivityDateTime, action.ActionPerformed, action.Amount, action.ProfileAction)
	if exeErr != nil {
		logger.Error("error execting the AddRecord exec", exeErr)
		return errors.NewInternalServerError("can't add record")
	}
	return nil
}
func (a *ActionRecorderRepository) RetrieveEmployeeActivityRecord(ctx context.Context, id string) (*model.EmployeeActivityActionRecorder, *errors.RestErr) {
	stmt, err := a.ActionDB.Prepare(queryEmployeeActivityRecorder)
	if err != nil {
		logger.Error("error preparing the QueryEmployeeRecord query", err)
		return nil, errors.NewInternalServerError("can not add record")
	}
	defer stmt.Close()

	var employee *model.EmployeeActivityActionRecorder
	if exeErr := stmt.QueryRow(id).Scan(employee.ActivityID, employee.EmployeeID, employee.ActivityDateTime, employee.ActionPerformed, employee.Amount, employee.ProfileAction); err != nil {
		logger.Error(fmt.Sprintf("error querying for activities of employee with ID: %s\n", id), exeErr)
		return nil, errors.NewInternalServerError("can't query customer activity")
	}

	return employee, nil
}
func (a *ActionRecorderRepository) AddTransactionActivityRecord(ctx context.Context, action *model.TransactionActivityActionRecorder) *errors.RestErr {
	//TODO: Add database transaction logic here
	return nil
}
func (a *ActionRecorderRepository) RetrieveTransactionActivityRecord(ctx context.Context, id string) (*model.TransactionActivityActionRecorder, *errors.RestErr) {
	//TODO: Add database transaction logic here
	return nil, nil
}
