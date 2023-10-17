package main

import (
	"bank_customers/src/domain/repository"
	"bank_customers/src/service"
	"bank_customers/src/utils/logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	"bank_customers/src/adapter/db"
	"bank_customers/src/adapter/http"

	"bank_customers/src/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_customers/src/utils/.env")

	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	var customerRepo repository.CustomerRepository
	//var accountCreatorInterface db.AccountCreationInterface

	datasourceName := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_schema"))
	//accountCreationConnection := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_schema"))

	conn := MysqlConnection(datasourceName)
	//accountConnection := MysqlConnection(accountCreationConnection)

	customerRepo = db.NewMySqlCustomerRepository(conn)
	//accountCreatorInterface = db.NewMySqlAccountCreationConn(accountConnection)

	customerService := service.NewCustomerService(customerRepo)
	customerHandler := http.NewCustomerHandler(customerService)

	//defer conn.Close()

	router.POST("/customers", middleware.RandomTokenMiddleware(), customerHandler.Create)
	router.GET("/customers/searchemail/:email", middleware.RandomTokenMiddleware(), customerHandler.GetCustomerByEmail)
	router.GET("/customers/searchfirstname/:firstname", middleware.RandomTokenMiddleware(), customerHandler.GetCustomerByFirstName)
	router.PATCH("/customers/searchemail/:email", middleware.RandomTokenMiddleware(), customerHandler.UpdateCustomerInfo)
	router.POST("/customers/login", customerHandler.LoginByEmail)

	router.Run()

}

func MysqlConnection(database string) *sql.DB {

	fmt.Println("Connecting to DB")
	var err error
	db, err := sql.Open("mysql", database)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	logger.Info("UserDB Database successful initiated")

	return db
}
