package main

import "fmt"

func SqlInjectionDemo() {
	email := Input("Enter email: ")
	city := Input("Enter city: ")

	// enter the following for email (as-is):
	// ' or true limit 1 -- 

	sql := fmt.Sprintf("select * from CUSTOMERS where email='%v' and city='%v'", email, city)

	fmt.Println(sql)
}
