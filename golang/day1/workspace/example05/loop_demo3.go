package main

import "fmt"

func menu() int {
	fmt.Println("*** Main Menu ***")
	fmt.Println("1. Add customer")
	fmt.Println("2. Search customer")
	fmt.Println("3. Edit customer")
	fmt.Println("4. List all customers")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func main() {
	/*
		ans := -1
		for ans != 5 {
			ans = menu()
			if ans != 5 {
				fmt.Println("Your choice is", ans)
				fmt.Printf("Choice %d is not implemented yet!\n", ans)
			}
		}
	*/
	for {
		ans := menu()
		if ans == 5 {
			break
		}
		fmt.Println("Your choice is", ans)
		fmt.Printf("Choice %d is not implemented yet!\n", ans)
	}
	println("Bye.")
}
