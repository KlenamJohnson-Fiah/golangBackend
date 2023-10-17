package jwt_verify

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		//fmt.Println(strArr[1])
		return strArr[1]
	}

	return " "
}

func EmployeeVerifyToken(r *http.Request) (*jwt.Token, error) {
	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("E_ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := EmployeeVerifyToken(r)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}

	return nil
}

type AccessDetails struct {
	AccessUuid string
	UserId     string
	UserRole   string
}

func EmployeeExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := EmployeeVerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, err
		}
		UserRole, ok := claims["role"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
			UserRole:   UserRole,
		}, nil
	}
	return nil, err
}
