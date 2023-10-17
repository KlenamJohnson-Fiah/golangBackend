package service

import (
	"bank_JWTauth/src/domain/model"
	"fmt"
	"log"
	"os"
	"time"

	"bank_JWTauth/src/utils/errors"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
)

const (
	customer = "customer"
	employee = "employee"
)

func CustomerCreateToken(clientID string) (*model.TokenDetails, *errors.RestErr) {
	td := model.TokenDetails{}

	td.AccessTokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessTokenDetails.AccessUuid = uuid.NewV4().String()

	td.RefreshTokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshTokenDetails.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessTokenDetails.AccessUuid
	atClaims["user_id"] = clientID
	atClaims["exp"] = td.AccessTokenDetails.AtExpires
	atClaims["role"] = customer

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)

	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	var err error

	td.AccessTokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("C_ACCESS_SECRET")))

	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Couldn't sign access token")
	}

	rtClaims := jwt.MapClaims{}

	rtClaims["authorized"] = true
	rtClaims["access_uuid"] = td.RefreshTokenDetails.RefreshUuid
	rtClaims["user_id"] = clientID
	rtClaims["exp"] = td.RefreshTokenDetails.RtExpires
	rtClaims["role"] = employee

	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.RefreshTokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("C_REFRESH_SECRET")))
	if err != nil {
		return nil, errors.NewInternalServerError("Couldn't sign refresh token")
	}

	return &td, nil

}

func EmployeeCreateToken(clientID string) (*model.TokenDetails, *errors.RestErr) {
	td := model.TokenDetails{}

	td.AccessTokenDetails.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessTokenDetails.AccessUuid = uuid.NewV4().String()

	td.RefreshTokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshTokenDetails.RefreshUuid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessTokenDetails.AccessUuid
	atClaims["user_id"] = clientID
	atClaims["exp"] = td.AccessTokenDetails.AtExpires
	atClaims["role"] = employee

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)

	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	var err error

	td.AccessTokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("E_ACCESS_SECRET")))

	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Couldn't sign access token")
	}

	rtClaims := jwt.MapClaims{}

	rtClaims["authorized"] = true
	rtClaims["access_uuid"] = td.RefreshTokenDetails.RefreshUuid
	rtClaims["user_id"] = clientID
	rtClaims["exp"] = td.RefreshTokenDetails.RtExpires
	rtClaims["role"] = employee

	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	td.RefreshTokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("E_REFRESH_SECRET")))
	if err != nil {
		return nil, errors.NewInternalServerError("Couldn't sign refresh token")
	}

	return &td, nil

}
