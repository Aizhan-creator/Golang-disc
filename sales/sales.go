package sales

type Sales struct {
	position string
	salary   int
	address  string
}

func NewSales(position string, salary int, address string) Sales {
	return Sales{
		position: position,
		salary:   salary,
		address:  address,
	}
}
func (s *Sales) SetPosition(position string) {
	s.position = position
}
func (s *Sales) SetSalary(salary int) {
	s.salary = salary
}
func (s *Sales) SetAddress(address string) {
	s.address = address
}
func (s *Sales) EmployeePosition() string {
	return s.position
}
func (s *Sales) EmployeeSalary() int {
	return s.salary
}
func (s *Sales) EmployeeAddress() string {
	return s.address
}
