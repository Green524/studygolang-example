package main

import (
	"time"
)

type Employee struct {
	ID       int
	Name     string
	Address  string
	Dob      time.Time
	Position string
	Salary   int
	ManageID int
}

var dilbert Employee

func EmployeeByID(id int) *Employee {
	return &dilbert
}
