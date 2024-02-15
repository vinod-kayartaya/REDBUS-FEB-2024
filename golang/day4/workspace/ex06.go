package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	sentences := []string{
		"go is a simple but powerful language",
		"go language combines the best things from c++, java and python",
		"my name is vinod and i live in bangalore",
	}
	words := []string{}
	var mx sync.Mutex

	for _, sentence := range sentences {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			parts := strings.Split(s, " ")
			mx.Lock()
			for _, word := range parts {
				words = append(words, word)
				time.Sleep(1 * time.Millisecond)
			}
			mx.Unlock()
			// there may be any code here which may be executed concurrently or parallely with
			// other goroutines
		}(sentence)
	}

	wg.Wait()
	fmt.Println(words)

}
