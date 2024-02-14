package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name    string
	Age     int
	Contact
}

type Contact struct {
	Email string
	Phone string
}


func (p Person) Print() {
	p.Name = strings.ToUpper(p.Name)
	fmt.Printf("Name   : %v\n", p.Name)
	fmt.Printf("Email  : %v\n", p.Email)
	fmt.Printf("Phone  : %v\n", p.Phone)
	fmt.Printf("Age    : %v\n", p.Age)
	fmt.Println()
}

func main() {
	var p1 Person
	p2 := Person{Name: "Ramesh", Age: 33}

	p2.Email = "ramesh@xmpl.com"
	p2.Phone = "9876509762"
	

	ptr := &p1

	(*ptr).Name = "Harish"
	(*ptr).Age = 12
	(*ptr).Contact.Email = "harish@xmpl.com"
	(*ptr).Contact.Phone = "5672822112"

	p1.Print()
	p2.Print()

	fmt.Printf("p1 is %v\n", p1)
	fmt.Printf("p2 is %v\n", p2)

}
