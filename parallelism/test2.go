package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello()
	
	wg.Add(1)
	go func() {
		fmt.Println("world!")
		wg.Done()
	}()
	
	wg.Wait()
}

func sayHello() {
	fmt.Println("helo")
	wg.Done()

}
 