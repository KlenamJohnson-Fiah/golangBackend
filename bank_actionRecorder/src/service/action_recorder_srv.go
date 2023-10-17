package service

import (
	"bank_actionRecorder/src/domain/model"
	"bank_actionRecorder/src/domain/repository"
	"bank_actionRecorder/src/utils/errors"
	"context"
)

type ActionRecorderServiceInterface interface {
	AddCustomerActivityRecord(context.Context, *model.CustomerActivityActionRecorder) *errors.RestErr
	RetrieveCustomerActivityRecord(context.Context, string) (*model.CustomerActivityActionRecorder, *errors.RestErr)
	AddEmployeeActivityRecord(context.Context, *model.EmployeeActivityActionRecorder) *errors.RestErr
	RetrieveEmployeeActivityRecord(context.Context, string) (*model.EmployeeActivityActionRecorder, *errors.RestErr)
}

type actionRecoderServices struct {
	actionsRecoderRepo repository.ActionRecorderRepository
}

func NewActionRecorderService(actionsRecoderRepo repository.ActionRecorderRepository) ActionRecorderServiceInterface {
	return &actionRecoderServices{
		actionsRecoderRepo: actionsRecoderRepo,
	}
}

func (a *actionRecoderServices) AddCustomerActivityRecord(ctx context.Context, action *model.CustomerActivityActionRecorder) *errors.RestErr {

	if err := a.actionsRecoderRepo.AddCustomerActivityRecord(ctx, action); err != nil {
		return err
	}

	return nil
}

func (a *actionRecoderServices) RetrieveCustomerActivityRecord(ctx context.Context, id string) (*model.CustomerActivityActionRecorder, *errors.RestErr) {

	customerActivity, err := a.actionsRecoderRepo.RetrieveCustomerActivityRecord(ctx, id)
	if err != nil {
		return nil, err
	}
	return customerActivity, nil
}

func (a *actionRecoderServices) AddEmployeeActivityRecord(ctx context.Context, action *model.EmployeeActivityActionRecorder) *errors.RestErr {
	if err := a.actionsRecoderRepo.AddEmployeeActivityRecord(ctx, action); err != nil {
		return err
	}
	return nil
}

func (a *actionRecoderServices) RetrieveEmployeeActivityRecord(ctx context.Context, id string) (*model.EmployeeActivityActionRecorder, *errors.RestErr) {

	employeeActivity, err := a.actionsRecoderRepo.RetrieveEmployeeActivityRecord(ctx, id)
	if err != nil {
		return nil, err
	}
	return employeeActivity, nil

}
