package main

import (
	"encapdemo/hr"
	"fmt"
)

func main() {
	fmt.Println("Encapsulation demo")

	var e1 hr.Employee
	e1.SetId(123)
	e1.SetName("Raja")
	e1.SetSalary(22300)
	e1.Print()
}
