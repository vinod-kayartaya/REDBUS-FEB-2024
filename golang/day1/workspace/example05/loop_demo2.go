package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments supplied!")
		return
	}
	fmt.Println("Here are the command line arguments: ")

	for index, value := range os.Args {
		fmt.Printf("%d --> %v\n", index, value)
	}
}
