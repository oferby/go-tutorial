package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
)

var n = flag.Bool("n", false, "omit trailing new line.")
var sep = flag.String("s", " ", "default seperator.")

func main() {

	last := len(os.Args) - 1
	fmt.Printf("number of args: %v, last: %v\n", len(os.Args), last)
	for i := 0; i < last; i++ {
		fmt.Printf("%v ", os.Args[i])
	}

	fmt.Printf("%v\n", os.Args[last])

	fmt.Println("another option:")

	for _, arg := range os.Args {
		fmt.Printf("%v ", arg)
	}
	fmt.Printf("\n")

	fmt.Println("yet another option:")
	fmt.Println(strings.Join(os.Args, " "))

	fmt.Println("with Flag:")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	} 

}
