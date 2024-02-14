package hr

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float64
}

func (e Employee) Print() {
	fmt.Printf("ID    : %v\n", e.id)
	fmt.Printf("Name  : %v\n", e.name)
	fmt.Printf("Salary: %v\n", e.salary)
}

func (e *Employee) SetId(id int) {
	if id < 0 || id > 999999 {
		panic("Invalid id. Must be between 0 and 999999")
	}
	e.id = id
}

func (e *Employee) SetName(name string) {
	if len(name) < 3 || len(name) > 25 {
		panic("Invalid name. Too small or too big!")
	}
	e.name = name
}

func (e *Employee) SetSalary(salary float64) {
	if salary < 20000 || salary > 999999 {
		panic("Invalid salary. Must be between 20000 and 999999")
	}
	e.salary = salary
}
