package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/Users/vinod/Desktop"

	fmt.Println()
	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		// var fileOrDir string
		// if entry.IsDir() {
		// 	fileOrDir = "dir"
		// } else {
		// 	fileOrDir = "file"
		// }

		// fmt.Printf("[%v] %v\n", fileOrDir, entry.Name())
		if !entry.IsDir() {
			fmt.Printf("%v\n", entry.Name())
		}
	}
}
