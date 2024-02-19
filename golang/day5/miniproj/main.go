package main

import (
	"fmt"
	"miniproj/dao"
	"miniproj/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	city := utils.Input("Enter city: ")
	for _, c := range dao.GetAllCustomersFromCity(city) {
		fmt.Println(c)
	}
}
