package administration

type Admin struct {
	position string
	salary   int
	address  string
}

func NewAdmin(position string, salary int, address string) Admin {
	return Admin{
		position: position,
		salary:   salary,
		address:  address,
	}
}
func (a *Admin) SetPosition(position string) {
	a.position = position
}
func (a *Admin) SetSalary(salary int) {
	a.salary = salary
}
func (a *Admin) SetAddress(address string) {
	a.address = address
}
func (a *Admin) EmployeePosition() string {
	return a.position
}
func (a *Admin) EmployeeSalary() int {
	return a.salary
}
func (a *Admin) EmployeeAddress() string {
	return a.address
}
