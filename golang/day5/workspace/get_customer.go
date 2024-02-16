package main

import (
	"fmt"
	"strconv"
)

func GetOneCustomer() {
	id, err := strconv.Atoi(Input("Enter id to search: "))
	checkForError(err)

	db := getDb()
	defer db.Close()
	
	sql := "SELECT NAME, CITY, EMAIL FROM CUSTOMERS WHERE ID=?"
	stmt, err := db.Prepare(sql)
	checkForError(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var name, city, email string
	if err = row.Scan(&name, &city, &email); err != nil {
		fmt.Println("No data found for id", id)
	} else {
		fmt.Printf("Name   : %s\n", name)
		fmt.Printf("City   : %s\n", city)
		fmt.Printf("Email  : %s\n", email)
		fmt.Println()
	}
}
