package repository

import (
	"bank_actionRecorder/src/domain/model"
	"bank_actionRecorder/src/utils/errors"
	"context"
)

type ActionRecorderRepository interface {
	AddCustomerActivityRecord(context.Context, *model.CustomerActivityActionRecorder) *errors.RestErr
	RetrieveCustomerActivityRecord(context.Context, string) (*model.CustomerActivityActionRecorder, *errors.RestErr)
	AddEmployeeActivityRecord(context.Context, *model.EmployeeActivityActionRecorder) *errors.RestErr
	RetrieveEmployeeActivityRecord(context.Context, string) (*model.EmployeeActivityActionRecorder, *errors.RestErr)
	AddTransactionActivityRecord(context.Context, *model.TransactionActivityActionRecorder) *errors.RestErr
	RetrieveTransactionActivityRecord(context.Context, string) (*model.TransactionActivityActionRecorder, *errors.RestErr)
}
