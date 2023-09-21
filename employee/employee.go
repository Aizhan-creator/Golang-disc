package employee

type Employee interface {
	EmployeePosition() string
	EmployeeSalary() int
	EmployeeAddress() string
}
