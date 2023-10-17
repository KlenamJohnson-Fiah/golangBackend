package app

import (
	"bank_actionRecorder/src/utils/logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	"bank_actionRecorder/src/adapter/framework_left/http"
	"bank_actionRecorder/src/domain/repository"
	"bank_actionRecorder/src/service"

	"bank_actionRecorder/src/adapter/framework_right/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// var router = gin.Default()

func componentInitialization() {

	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_customers/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	var actionRecorderRepo repository.ActionRecorderRepository
	datasourceActionRecorder := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), "record")
	connActionRecorder := MysqlConnection(datasourceActionRecorder)
	actionRecorderRepo = db.NewActionRecorderRepository(connActionRecorder)
	actionRecorderService := service.NewActionRecorderService(actionRecorderRepo)
	actionRecorderHandler := http.NewActionRecorderHandler(actionRecorderService)

	router.POST("/actions/customer/record", actionRecorderHandler.AddCustomerActionHandler)
	router.GET("/actions/customer/retrieve", actionRecorderHandler.RetrieveCustomerActionHandler)
	router.POST("/actions/employee/record", actionRecorderHandler.AddEmployeeActionHandler)
	router.GET("/actions/employee/retrieve", actionRecorderHandler.RetrieveEmployeeActionHandler)
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
