package service

import (
	"bank_JWTauth/src/domain/model"
	"bank_JWTauth/src/domain/repository"
	"bank_JWTauth/src/utils/errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type AccessTokenInterface interface {
	CustomerLogin(email string, password string) (*model.TokenDetails, *errors.RestErr)
	EmployeeLogin(email string, password string) (*model.TokenDetails, *errors.RestErr)
}

type accessTokenService struct {
	accessTokenRepo repository.AccessTokenRepository
}

func NewAccessTokenService(accessTokenRepo repository.AccessTokenRepository) AccessTokenInterface {
	return &accessTokenService{
		accessTokenRepo: accessTokenRepo,
	}
}

var (
	client = resty.New()
)

func (a *accessTokenService) CustomerLogin(email string, password string) (*model.TokenDetails, *errors.RestErr) {
	//TODO: Figure out how to use resty,Reponse and Bind it to a defined struct
	//TODO: Fix errors that from other microservices call with resty
	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := client.R().SetBody(model.Login{Email: email, Password: password}).SetAuthToken(os.Getenv("AUTH_TOKEN")).Post("http://localhost:8080/customers/login")
	fmt.Println(resp.Status())
	switch {
	case err != nil:
		return nil, errors.NewInternalServerError("Error in response")
	case strings.Contains(resp.String(), "error"):
		return nil, errors.NewInternalServerError("Error in response")
	case resp.Status() >= "299":
		return nil, errors.NewInternalServerError("Error in response")
	}

	customerUUIDExtractor := resp.String()[7:43]

	tokenMaker, helpErr := CustomerCreateToken(customerUUIDExtractor)
	if helpErr != nil {
		fmt.Println(helpErr)
		return nil, errors.NewBadRequestError("error creatin token")
	}

	if tokenSaveErr := a.accessTokenRepo.SaveToken(tokenMaker); tokenSaveErr != nil {
		return nil, tokenSaveErr
	}

	return tokenMaker, nil

}

func (a *accessTokenService) EmployeeLogin(email string, password string) (*model.TokenDetails, *errors.RestErr) {
	//TODO: Figure out how to use resty,Reponse and Bind it to a defined struct
	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := client.R().SetBody(model.Login{Email: email, Password: password}).SetAuthToken(os.Getenv("AUTH_TOKEN")).Post("http://localhost:8084/employee/login")
	if err != nil {
		log.Println(err)
		return nil, errors.NewBadRequestError("Error on relay")
	}
	fmt.Println(resp.Status())
	fmt.Println(resp)

	var EmployeeUUIDExtractor string
	if resp.StatusCode() >= 299 {
		return nil, errors.NewInternalServerError("Error in response")
	}

	EmployeeUUIDExtractor = resp.String()[7:43]
	fmt.Println(EmployeeUUIDExtractor)

	tokenMaker, helpErr := EmployeeCreateToken(EmployeeUUIDExtractor)
	if helpErr != nil {
		fmt.Println(helpErr)
		return nil, errors.NewBadRequestError("error creatin token")
	}

	if tokenSaveErr := a.accessTokenRepo.SaveToken(tokenMaker); tokenSaveErr != nil {
		return nil, tokenSaveErr
	}

	return tokenMaker, nil

}
