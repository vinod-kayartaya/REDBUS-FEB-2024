package main

import "fmt"

func AddCustomerData() {
	db := getDb()
	defer db.Close()
	sql := "INSERT INTO CUSTOMERS (NAME, CITY, EMAIL) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(sql) // first roundtrip to db; server will keep a compiled copy of this
	checkForError(err)
	defer stmt.Close()
	for {
		fmt.Println("Enter customer details")
		name := Input("Name : ")
		city := Input("City : ")
		email := Input("Email: ")
		result, err := stmt.Exec(name, city, email) // subsequent roundtrips to db; server executes 
													// the compiled sql with these values
		checkForError(err)
		newId, err := result.LastInsertId()
		checkForError(err)
		fmt.Printf("New customer data added successfully with ID %v\n", newId)
		choice := Input("Want to add another? (yes/no) [yes] ")
		if choice == "no" {
			break
		}
	}
}

func AcceptAndAddCustomerData() {
	db := getDb()
	defer db.Close()

	for {
		fmt.Println("Enter customer details")
		name := Input("Name : ")
		city := Input("City : ")
		email := Input("Email: ")

		sql := fmt.Sprintf("INSERT INTO CUSTOMERS(NAME, CITY, EMAIL) VALUES('%v', '%v', '%v')",
			name, city, email)

		_, err := db.Exec(sql)
		checkForError(err)

		fmt.Println("New customer data added successfully.")
		choice := Input("Want to add another? (yes/no) [yes] ")
		if choice == "no" {
			break
		}
	}
}
