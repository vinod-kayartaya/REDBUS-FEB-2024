package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(<-ch)
		// ch <- "kumar"
		fmt.Println("5 second gr is done")
	}()

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("10 second gr is done")
		fmt.Println(<-ch)
		// ch <- "kumar"
	}()



	
	ch <- "vinod-1"
	ch <- "vinod-2"
	ch <- "vinod-3"
	ch <- "vinod-4"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch <- "shyam"
	ch <- "harish"

	fmt.Println(<-ch)
	fmt.Scanln()

}
