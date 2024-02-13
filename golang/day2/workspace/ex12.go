package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		content := scanner.Text()
		fmt.Printf("%2d. %s\n", i, content)
	}
}
