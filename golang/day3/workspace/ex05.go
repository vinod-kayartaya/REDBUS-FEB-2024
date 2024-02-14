package main

import "fmt"

type Product struct {
	Id    int
	Name  string
	Price float64
}

func (p *Product) Print() {
	p.Price += 500
	fmt.Printf("ID     : %d\n", p.Id)
	fmt.Printf("Name   : %v\n", p.Name)
	fmt.Printf("Price  : %v\n", p.Price)
	fmt.Println()
}

func main() {
	p1 := Product{Name: "Logitech Mouse", Price: 4500, Id: 23}
	p2 := Product{Name: "Apple Magic Mouse", Price: 8500, Id: 28}
	var p3 Product
	p1.Print()
	p2.Print()
	p3.Print()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
