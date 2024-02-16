package dao

import (
	"miniproj/model"
	"miniproj/utils"
)

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
