package http

import (
	"bank_JWTauth/src/domain/model"
	"bank_JWTauth/src/service"
	"bank_JWTauth/src/utils/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	CustomerLogin(c *gin.Context)
	EmployeeLogin(c *gin.Context)
}

type accesstokenHandler struct {
	accessTokenService service.AccessTokenInterface
}

func NewAccessTokenHandler(accessTokenService service.AccessTokenInterface) AccessTokenHandler {
	return &accesstokenHandler{
		accessTokenService: accessTokenService,
	}
}

func (a *accesstokenHandler) CustomerLogin(c *gin.Context) {
	//TODO: Find Binding(JSON and form-data struggle)
	var loginRequest model.Login
	if err := c.ShouldBind(&loginRequest); err != nil {
		reqErr := errors.NewBadRequestError("invalid json")
		c.JSON(reqErr.Status, err)
		return

	}
	fmt.Println(loginRequest)

	customer, loginErr := a.accessTokenService.CustomerLogin(loginRequest.Email, loginRequest.Password)
	if loginErr != nil {
		c.JSON(loginErr.Status, loginErr)
		return

	}

	c.JSON(http.StatusFound, customer)
	return

}

func (a *accesstokenHandler) EmployeeLogin(c *gin.Context) {
	//TODO: Find Binding(JSON and form-data struggle)
	var loginRequest model.Login
	if err := c.ShouldBind(&loginRequest); err != nil {
		reqErr := errors.NewBadRequestError("invalid json")
		c.JSON(reqErr.Status, err)
		return

	}
	fmt.Println(loginRequest)

	customer, loginErr := a.accessTokenService.EmployeeLogin(loginRequest.Email, loginRequest.Password)
	if loginErr != nil {
		c.JSON(loginErr.Status, loginErr)
		return

	}

	c.JSON(http.StatusFound, customer)
	return

}
