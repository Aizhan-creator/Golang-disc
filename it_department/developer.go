package it_department

type Developer struct {
	Position string
	Salary   int
	Address  string
}

func (d *Developer) EmployeePosition() string {
	return d.Position
}

func (d *Developer) EmployeeSalary() int {
	return d.Salary
}

func (d *Developer) EmployeeAddress() string {
	return d.Address
}
