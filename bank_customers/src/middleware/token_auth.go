package middleware

import (
	jwtVerify "bank_customers/src/utils/jwt_verify"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RandomTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
		if environ != nil {
			log.Fatalf("Error loading .env file")
		}
		metaData, err := jwtVerify.EmployeeExtractTokenMetadata(c.Request)
		if err != nil {
			//fmt.Println(err)
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Set("access_uuid", metaData.AccessUuid)
		c.Set("user_id", metaData.UserId)
		c.Set("role", metaData.UserRole)
		c.Next()
	}
}
