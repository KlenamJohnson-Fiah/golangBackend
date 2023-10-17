package repository

import (
	"bank_JWTauth/src/domain/model"
	"bank_JWTauth/src/utils/errors"
)

type AccessTokenRepository interface {
	SaveToken(accessToken *model.TokenDetails) *errors.RestErr
	DeleteAccessToken(uuid string) *errors.RestErr
}
