package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int, 10)
	// empty struct hold no memory
	doneCh := make(chan struct{})
	
	wg.Add(2)
	
	go func() {
loopMark:
		for {
			select {
			case i := <-ch:
				fmt.Printf("got %v\n", i)
			case <-doneCh:
				fmt.Println("got end signal")
				break loopMark
			}
		}
		
		fmt.Println("Done!")
		wg.Done()

	}()

	go func(ch chan<- int) {
		ch <- 42
		ch <- 22

		time.Sleep(100 * time.Millisecond)
		doneCh <- struct{}{}
		wg.Done()
	}(ch)

	wg.Wait()

}

