package main

import "fmt"

func main(){
	var d, m, y int
	fmt.Print("Enter date in d/m/y format: ")
	fmt.Scanf("%d/%d/%d", &d, &m, &y)
	fmt.Printf("You entered %d-%d-%d\n", y, m, d)
}