 package db

import (
	cassandra "bookstore_oauth-api/src/clients/cassandra_client"
	"bookstore_oauth-api/src/domain/access_token"
	errors "bookstore_oauth-api/src/utils/errors_utils"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(*access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(*access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (d *dbRepository) Create(at *access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserID, at.ClientID, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}

func (d *dbRepository) UpdateExpirationTime(at *access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).Scan(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil

}
