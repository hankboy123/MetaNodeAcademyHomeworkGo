package main

import (
	"fmt"
)

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("employee", e.Age, e.Name, e.EmployeeID)
}
