package model

import "time"

type Customers struct {
	Id               string    `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	Status           string    `json:"status"`
	Password         string    `json:"password"`
	DateCreated      time.Time `json:"date_created"`
	DateLastModified time.Time `json:"date_last_modified"`
}

type MultipleCustomers []Customers

type CustomerDetailsOutput struct {
	Id               string    `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	Status           string    `json:"status"`
	DateCreated      time.Time `json:"date_created"`
	DateLastModified time.Time `json:"date_last_modified"`
}

type MultipleCustomersDetailsOutput []CustomerDetailsOutput

type CustomerDetailsPatch struct {
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Address          string    `json:"address"`
	DateLastModified time.Time `json:"date_last_modified"`
}

type CustomerLoginCredentials struct {
	//Id       string `json:"id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password"  json:"password"`
}

type CustomerLoginSuccessful struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
