package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	componentInitialization()

	router.Run(":8081")
}

// import (
// 	//"bank_customers/src/adapter/db"
// 	"bank_actionRecorder/src/utils/logger"
// 	"database/sql"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	router = gin.Default()
// )

// func MysqlConnection(database string) *sql.DB {

// 	fmt.Println("Connecting to DB")
// 	var err error
// 	db, err := sql.Open("mysql", database)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}

// 	logger.Info("UserDB Database successful initiated")

// 	return db
// }

// func StartApplication() {
// 	Mapurls()

// 	logger.Info("about to start the application")
// 	router.Run(":8080")

// }
