package main

import "fmt"

func main() {
	var month int
	fmt.Print("Enter a month number (1-12): ")
	fmt.Scan(&month)

	var days int

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		var year int
		fmt.Print("Enter the year: ")
		fmt.Scan(&year)

		if year <= 0 {
			fmt.Println("Invalid year.")
			return
		}
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			days = 29
		} else {
			days = 28
		}
	default:
		fmt.Println("Invalid month.")
		return
	}

	fmt.Printf("Month %d has %d days\n", month, days)

}
