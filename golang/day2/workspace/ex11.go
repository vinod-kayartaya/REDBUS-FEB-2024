package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var filename string

	fmt.Print("Enter filename to read: ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, _ := io.ReadAll(file)
	fmt.Print(string(content))

}
