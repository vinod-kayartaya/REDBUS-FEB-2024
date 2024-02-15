package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type CustomerDAO struct {
	Customers []Customer `json:"customers"`
}

func (dao *CustomerDAO) LoadFromJsonFile() {
	// 1. read the name of the file from the user
	filename := Input("Enter JSON source filename: ")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Couldn't open the file. Check logs.")
		return
	}
	defer file.Close()
	jsonContent, _ := io.ReadAll(file)

	// 2. unmarshal the JSON content into `Customers` slice
	if err := json.Unmarshal([]byte(jsonContent), &dao.Customers); err != nil {
		fmt.Println("There was an error while converting JSON data into Customers")
		return
	}

	fmt.Printf("Loaded %d customers from `%v` file\n", len(dao.Customers), filename)
}

func (dao CustomerDAO) GetAllCustomers() []Customer {
	return dao.Customers
}
