package middleware

import (
	jwtVerify "bank_employee/src/utils/jwt_verify"
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type ctxKey string

const (
	access_uuid ctxKey = "access_uuid"
	user_id     ctxKey = "user_id"
	role        ctxKey = "role"
)

func SetContextValue(ctx context.Context, key ctxKey, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextValue(ctx context.Context, key ctxKey) string {
	ctxValue := ctx.Value(key).(string)
	return ctxValue

}

func RandomTokenMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_JWTauth/src/utils/.env")
		if environ != nil {
			log.Fatalf("Error loading .env file")
		}
		metaData, err := jwtVerify.EmployeeExtractTokenMetadata(r)
		if err != nil {
			//fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			//JSON(http.StatusUnauthorized, "unauthorized")
			//Abort()
			return

		}

		ctx = SetContextValue(ctx, access_uuid, metaData.AccessUuid)
		ctx = SetContextValue(ctx, user_id, metaData.UserId)
		ctx = SetContextValue(ctx, role, metaData.UserRole)
		r = r.WithContext(ctx)

		//fmt.Println(ctx.Value(access_uuid))

		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)

}
