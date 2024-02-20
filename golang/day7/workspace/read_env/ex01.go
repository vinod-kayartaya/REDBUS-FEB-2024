package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	fmt.Println("------------------------------------------------")

	path := os.Getenv("PATH")
	vinod_email := os.Getenv("VINOD_EMAIL")
	vinod_phone := os.Getenv("VINOD_PHONE")

	fmt.Println("path is", path)
	fmt.Println("vinod_email is", vinod_email)
	if vinod_phone == "" {
		fmt.Println("No phone found for Vinod")
	} else {
		fmt.Println("vinod_phone is", vinod_phone)
	}

}
