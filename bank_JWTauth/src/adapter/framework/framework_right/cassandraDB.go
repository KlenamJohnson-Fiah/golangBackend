package frameworkright

import (
	"bank_JWTauth/src/domain/repository"

	"bank_JWTauth/src/domain/model"

	"bank_JWTauth/src/utils/errors"
	"bank_JWTauth/src/utils/jwt_verify"

	"bank_JWTauth/src/utils/logger"

	"github.com/gocql/gocql"
)

type cassandra struct {
	cassandraDB *gocql.Session
}

func NewCassandraRepository(cassandraDB *gocql.Session) repository.AccessTokenRepository {
	return &cassandra{
		cassandraDB: cassandraDB,
	}

}

func (c *cassandra) SaveToken(accessToken *model.TokenDetails) *errors.RestErr {

	if err := c.cassandraDB.Query(queryToAddAccessToken,
		accessToken.AccessTokenDetails.AccessUuid,
		accessToken.AccessTokenDetails.AccessToken,
		accessToken.AccessTokenDetails.AtExpires).Exec(); err != nil {
		logger.Error("Error creating access_token", err)
		return errors.NewInternalServerError("couldn't create access token in DB")
	}
	if err := c.cassandraDB.Query(queryToAddRefreshToken,
		accessToken.RefreshTokenDetails.RefreshUuid,
		accessToken.RefreshTokenDetails.RefreshToken,
		accessToken.RefreshTokenDetails.RtExpires,
	).Exec(); err != nil {
		logger.Error("Error creating refresh_token", err)
		return errors.NewInternalServerError("couldn't create refresh token in DB")
	}

	return nil
}

func (c *cassandra) DeleteAccessToken(uuid string) *errors.RestErr {

	if err := c.cassandraDB.Query(
		queryToDeleteAccessToken,
		uuid,
	).Exec(); err != nil {
		logger.Error("can delete Token", err)
		return errors.NewInternalServerError("can delete Token")
	}
	return nil
}
func (c *cassandra) DeleteRefreshToken(uuid string) *errors.RestErr {

	if err := c.cassandraDB.Query(
		queryToDeleteRefreshToken,
		uuid,
	).Exec(); err != nil {
		logger.Error("can delete Token", err)
		return errors.NewInternalServerError("can delete Token")
	}
	return nil
}

func (c *cassandra) FetchAuth(authD *jwt_verify.AccessDetails) (string, error) {
	err := c.cassandraDB.Query(QueryToVerifyAccessTokenUUID, authD.AccessUuid).Exec()
	if err != nil {
		return "", err
	}

	return authD.AccessUuid, nil
}
