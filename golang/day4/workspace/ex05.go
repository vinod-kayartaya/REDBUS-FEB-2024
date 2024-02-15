package main

import (
	"fmt"
	"time"
)

func gr1(ch chan string) {
	for {
		ch <- "channel1"
		time.Sleep(500 * time.Millisecond)
	}
}

func gr2(ch chan string) {
	for {
		ch <- "channel2"
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)

	fmt.Println("starting gr1..")
	go gr1(ch1)
	fmt.Println("starting gr2..")
	go gr2(ch2)

	for {
		select {
		case channel1Data := <-ch1:
			fmt.Println("Got data from", channel1Data)
		case channel2Data := <-ch2:
			fmt.Println("Got data from", channel2Data)
		}
	}

}
