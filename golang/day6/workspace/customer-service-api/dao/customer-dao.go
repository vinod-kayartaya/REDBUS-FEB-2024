package dao

import (
	"api/model"
	"api/utils"
)

func AddCustomer(customer model.Customer) int {
	db := connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO CUSTOMERS(NAME, CITY, EMAIL) VALUES(?, ?, ?)")
	utils.CheckForError(err)
	defer stmt.Close()

	result, err := stmt.Exec(customer.Name, customer.City, customer.Email)
	utils.CheckForError(err)

	newId, _ := result.LastInsertId()
	return int(newId)
}

func GetOneCustomer(id int) *model.Customer {

	db := connect()
	stmt, err := db.Prepare("select ID, NAME, CITY, EMAIL from CUSTOMERS where ID=?")
	utils.CheckForError(err)

	row := stmt.QueryRow(id)
	utils.CheckForError(err)

	var c model.Customer
	err = row.Scan(&c.Id, &c.Name, &c.City, &c.Email)

	if err != nil {
		return nil
	}
	return &c
}

func GetAllCustomers() []model.Customer {

	db := connect()
	stmt, err := db.Prepare("select ID, NAME, CITY, EMAIL from CUSTOMERS")
	utils.CheckForError(err)

	rows, err := stmt.Query()
	utils.CheckForError(err)

	customers := []model.Customer{}

	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Email)
		utils.CheckForError(err)
		customers = append(customers, c)
	}

	return customers

}

func GetAllCustomersFromCity(city string) []model.Customer {

	db := connect()
	stmt, err := db.Prepare("select ID, NAME, CITY, EMAIL from CUSTOMERS where CITY=?")
	utils.CheckForError(err)

	rows, err := stmt.Query(city)
	utils.CheckForError(err)

	customers := []model.Customer{}

	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Email)
		utils.CheckForError(err)
		customers = append(customers, c)
	}

	return customers

}
