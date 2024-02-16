package main

import (
	"fmt"
	"miniproj/dao"
)

func main() {

	for _, c := range dao.GetAllCustomers() {
		fmt.Println(c)
	}
}
