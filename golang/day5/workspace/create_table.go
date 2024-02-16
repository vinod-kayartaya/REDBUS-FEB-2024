package main

import "fmt"

func CreateTable() {
	db := getDb()

	sql := `CREATE TABLE CUSTOMERS (
		ID INTEGER PRIMARY KEY AUTO_INCREMENT,
		NAME varchar(50) NOT NULL,
		EMAIL varchar(50) UNIQUE,
		CITY varchar(50)
	)`

	_, err := db.Exec(sql)
	checkForError(err)
	fmt.Println("Table created successfully!")
}
