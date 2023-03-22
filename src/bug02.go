package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go Print(ch, &wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}