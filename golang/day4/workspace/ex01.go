package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doStuff(msg string, limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d. %v\n", i, msg)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	fmt.Println("start of main()")
	wg.Add(2)
	go doStuff("Hello, world!", 10)
	go doStuff("Welcome to `go` training!", 5)

	wg.Wait()
	fmt.Println("end of main()")
}
