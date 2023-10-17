package application

import (
	"bank_JWTauth/src/adapter/framework/framework_left/http"
	"bank_JWTauth/src/domain/repository"
	"bank_JWTauth/src/service"
	"fmt"

	frameworkright "bank_JWTauth/src/adapter/framework/framework_right"

	"github.com/gocql/gocql"
)

func ComponenentsInitialization() {

	var accessTokenRepo repository.AccessTokenRepository
	accessTokenRepo = frameworkright.NewCassandraRepository(CassandraConnection())
	accessTokenService := service.NewAccessTokenService(accessTokenRepo)
	accessTokenHandler := http.NewAccessTokenHandler(accessTokenService)

	router.GET("/accesstoken/customer/login", accessTokenHandler.CustomerLogin)
	router.GET("/accesstoken/employee/login", accessTokenHandler.EmployeeLogin)
}

func CassandraConnection() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tokens"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra setup done")
	return session

}
