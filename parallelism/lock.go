package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg = sync.WaitGroup{}
var lock = sync.RWMutex{}
var counter = 0 

func main() {

	fmt.Printf("Num of Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(20)

	for i := 0; i < 10; i++ {

		wg.Add(2)
		go printCounter()
		go increment()

	}
	
	wg.Wait()

	fmt.Println("another way...")

	// don't do that. Not async
	for i := 0; i < 10; i++ {

		wg.Add(2)
		lock.RLock()
		go printCounter2()
		lock.Lock()
		go increment2()

	}
	
	wg.Wait()

}



func printCounter() {
	lock.RLock()
	fmt.Println(counter)
	lock.RUnlock()
	wg.Done()

}

func increment() {
	lock.Lock()
	counter++;
	lock.Unlock()
	wg.Done()
}
 


func printCounter2() {
	
	fmt.Println(counter)
	lock.RUnlock()
	wg.Done()

}

func increment2() {
	
	counter++;
	lock.Unlock()
	wg.Done()
}
 