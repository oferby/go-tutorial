package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int, 10)

	wg.Add(2)
	
	go func(ch <-chan int) {

		for i := range ch {
			fmt.Printf("got %v\n", i)
		}

		wg.Done()

	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 22
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()

}

