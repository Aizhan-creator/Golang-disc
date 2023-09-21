package hr

type Hr struct {
	position string
	salary   int
	address  string
}

func NewHr(position string, salary int, address string) Hr {
	return Hr{
		position: position,
		salary:   salary,
		address:  address,
	}
}
func (h *Hr) SetPosition(position string) {
	h.position = position
}
func (h *Hr) SetSalary(salary int) {
	h.salary = salary
}
func (h *Hr) SetAddress(address string) {
	h.address = address
}
func (h *Hr) EmployeePosition() string {
	return h.position
}
func (h *Hr) EmployeeSalary() int {
	return h.salary
}
func (h *Hr) EmployeeAddress() string {
	return h.address
}
