package main

import (
	"fmt"
	"time"

)


func main() {
	go fmt.Println("helo")
	
	go func() {
		fmt.Println("world!")
	}()
	
	time.Sleep(100 * time.Millisecond)
}
 