package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string

	fmt.Print("Enter filename to create: ")
	fmt.Scan(&filename)

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var name string
		fmt.Print("Enter name (QUIT to stop): ")
		// fmt.Scanln(&name)
		scanner.Scan()
		name = scanner.Text()

		if name == "QUIT" {
			break
		}

		// file.WriteString(name)
		// file.WriteString("\n")
		fmt.Fprintln(file, name)
	}

	fmt.Println("End of main()")

}
