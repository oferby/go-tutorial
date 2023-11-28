package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)

	wg.Add(2)
	
	go func() {
		i := <-ch
		fmt.Printf("got %v\n", i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()

	wg.Wait()

}

