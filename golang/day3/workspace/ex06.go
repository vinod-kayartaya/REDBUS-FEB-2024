package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	GetName() string
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14167
}

func (t Circle) GetName() string {
	return "circle"
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) GetName() string {
	return "triangle"
}

// polymorphic method
func ProcessShape(c1 Shape) {
	fmt.Printf("Area of the %v is %v\n", c1.GetName(), c1.Area())
}

func main() {
	c1 := Circle{12.34}
	t1 := Triangle{1.2, 3.4}

	ProcessShape(c1)
	ProcessShape(t1)
}
