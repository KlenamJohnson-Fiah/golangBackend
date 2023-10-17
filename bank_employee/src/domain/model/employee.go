package model

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	Branch    string `json:"branch"`
}

type EmployeeLoginCredentials struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password"  json:"password"`
}

type EmployeeLoginSuccessful struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SliceOfEmployees []Employee
