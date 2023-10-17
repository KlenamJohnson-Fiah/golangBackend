package componentsinit

import (
	"bank_transaction/src/adapter/framework_left/http"
	"bank_transaction/src/adapter/framework_right/db"
	"bank_transaction/src/domain/repository"
	"bank_transaction/src/service"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ComponentInit() {
	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_customers/src/utils/.env")

	if environ != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	var transactionRepo repository.TransactionRepository
	datasourceName := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_schema"))
	DBconn := MySqlInit(datasourceName)
	transactionRepo = db.NewMySqlAccountRepository(DBconn)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := http.NewTransactionHandler(transactionService)

	router.GET("/account/check/:accountNumber", transactionHandler.CheckBalanceHandler)
	router.PATCH("/account/deposit", transactionHandler.DepositeHandler)
	router.PATCH("/account/withdraw", transactionHandler.WithdrawHandler)

	router.Run(":8086")
}

func MySqlInit(dbstring string) *sql.DB {
	fmt.Println("Connecting to DB")
	var err error
	db, err := sql.Open("mysql", dbstring)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("UserDB Database successful initiated")

	return db

}
