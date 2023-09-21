package finance

type Finance struct {
	position string
	salary   int
	address  string
}

func NewFinance(position string, salary int, address string) Finance {
	return Finance{
		position: position,
		salary:   salary,
		address:  address,
	}
}
func (f *Finance) SetPosition(position string) {
	f.position = position
}
func (f *Finance) SetSalary(salary int) {
	f.salary = salary
}
func (f *Finance) SetAddress(address string) {
	f.address = address
}
func (f *Finance) EmployeePosition() string {
	return f.position
}
func (f *Finance) EmployeeSalary() int {
	return f.salary
}
func (f *Finance) EmployeeAddress() string {
	return f.address
}
