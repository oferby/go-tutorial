package main

import (
	"fmt"
	"os"
	"bufio"
)

func printUsage() {
	fmt.Println("USAGE: command file-name")
}


func main() {
	files := os.Args[1:]

	if len(files) != 1 {
		printUsage()
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error in file name.")
		printUsage()
		os.Exit(1)
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Println(input.Text())


	}

	f.Close()	

}