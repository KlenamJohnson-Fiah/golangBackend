package application

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	frameworkhttp "bank_employee/src/adapter/framework/framework_left/http"
	"bank_employee/src/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	createEmployeeTable = "CREATE TABLE employees (id VARCHAR (60) PRIMARY KEY,first_name VARCHAR (30) NOT NULL,last_name VARCHAR (30) NOT NULL,email VARCHAR (30) UNIQUE NOT NULL,password VARCHAR (70) NOT NULL,role VARCHAR (10),branch VARCHAR (30),status VARCHAR (10));"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {

	host := "127.0.0.1"
	port := "5432"

	environ := godotenv.Load("/Users/klenam/Documents/go/src/bank_employee/src/utils/.env")
	if environ != nil {
		log.Fatalf("Error loading .env file")
	}
	pgConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("pgsql_username"), os.Getenv("pgsql_password"), host, port, "employees")

	pgConnection := PostgreSQLConnection(pgConn)

	router.Handle("/employee", frameworkhttp.Create(pgConnection)).Methods("POST") //frameworkhttp.Create(pgConnection)).Methods("POST")
	router.Handle("/employee/login", frameworkhttp.Login(pgConnection)).Methods("POST")
	router.Handle("/employee/search", middleware.RandomTokenMiddleware(frameworkhttp.GetByName(pgConnection))).Methods("GET")

	fmt.Println("listening ...")
	log.Fatal(http.ListenAndServe(":8084", router))

}

func PostgreSQLConnection(database string) *sql.DB {
	fmt.Println("Connecting to DB")
	var err error
	db, err := sql.Open("postgres", database)
	if err != nil {
		fmt.Println("1")
		panic(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("2")
		panic(err)
	}

	// _, err = db.Exec(createEmployeeTable)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("UserDB Database successful initiated")

	return db
}
