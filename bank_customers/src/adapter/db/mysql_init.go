package db

// import (
// 	"bank_customers/src/utils/logger"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// // var (
// // 	Client *sql.DB
// // )

// func init() {
// 	environ := godotenv.Load("bank_customers/src/utils/.env")

// 	if environ != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	datasourceName := fmt.Sprintf("%s:%s@tcp/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_schema"))

// 	var err error
// 	Client, err := sql.Open("mysql", datasourceName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = Client.Ping(); err != nil {
// 		panic(err)
// 	}
// 	log.Println("UserDB Database successful initiated")
// 	logger.Info("UserDB Database successful initiated")
// }
