package http

import (
	"bank_employee/src/domain/model"
	"bank_employee/src/middleware"
	"bank_employee/src/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EmployeeHandlerInterface interface {
	Create(pg *sql.DB) http.Handler
	LoginLogin(pg *sql.DB) http.Handler
	GetByName(pg *sql.DB) http.Handler
}

func Create(pg *sql.DB) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		var employee model.Employee

		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &employee)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

		newData, err := json.Marshal(employee)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(newData))

		employeeCreate, CreateErr := service.CreateEmployee(pg, &employee)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("couldn't create user"))
			fmt.Println(CreateErr)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		x, _ := json.Marshal(employeeCreate)
		w.Write(x)
		json.NewEncoder(w).Encode(employeeCreate)

	}
	return http.HandlerFunc(fn)

}

func Login(pg *sql.DB) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var employeeLoginCred model.EmployeeLoginCredentials

		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &employeeLoginCred)
		respnse, Lerr := service.Login(pg, &employeeLoginCred)
		if Lerr != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("user might not exist"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respnse)

	}
	return http.HandlerFunc(fn)
}

func GetByName(pg *sql.DB) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := r.URL.Query()
		filter := query.Get("name")

		employee, err := service.GetByName(pg, filter)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("user might not exist"))
			return
		}
		fmt.Println(employee)
		x, _ := json.Marshal(employee)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		w.Write(x)
		fmt.Println(middleware.GetContextValue(ctx, "role"))

	}
	return http.HandlerFunc(fn)
}
