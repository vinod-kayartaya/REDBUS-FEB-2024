package main

import (
	"fmt"
	"strconv"
)

func Menu() int {
	menuOptions := []string{
		"Exit",
		"Load data from JSON file",
		"Save data to JSON file",
		"Add new customer",
		"Update customer data",
		"Delete customer data",
		"Search by city",
		"List all customers",
	}
	for {
		fmt.Println("*** Main Menu ***")
		for index, option := range menuOptions {
			fmt.Printf("%d. %s\n", index, option)
		}

		choice, err := strconv.Atoi(Input("Enter your choice: "))
		if err != nil || choice < 0 || choice >= len(menuOptions) {
			fmt.Println("Invalid choice. Please retry")
		} else {
			return choice
		}
	}
}
