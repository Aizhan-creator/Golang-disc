package main

import (
	"fmt"
	"task1/administration"
	"task1/employee"
	"task1/finance"
	"task1/hr"
	"task1/it_department"
	"task1/sales"
)

func main() {
	sales := sales.NewSales("Sales Manager", 20000, "Astana")
	sales.SetPosition("Sales Manager")

	hr := hr.NewHr("HR Manager", 20000, "Astana")
	hr.SetPosition("HR Manager")

	finance := finance.NewFinance("Assistant", 20000, "Astana")
	finance.SetPosition("Assistant")

	admin := administration.NewAdmin("IT Admin", 20000, "Astana")
	admin.SetPosition("IT Admin")

	developer := it_department.Developer{Position: "Backend Developer"}

	ex1 := EmployeePosition(&sales)
	ex2 := EmployeePosition(&developer)
	ex3 := EmployeePosition(&finance)
	ex4 := EmployeePosition(&admin)
	ex5 := EmployeePosition(&hr)

	fmt.Println("Position for Sales department: " + ex1)
	fmt.Println("Position for IT department: " + ex2)
	fmt.Println("Position for Finance department: " + ex3)
	fmt.Println("Position for Administration department: " + ex4)
	fmt.Println("Position for HR department: " + ex5)
	//dck.SecretFunction()
}

func EmployeePosition(e employee.Employee) string {
	return e.EmployeePosition()
}
