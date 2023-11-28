package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main()  {
	
	if len(os.Args) != 2 {
		fmt.Println("USAGE: readfile <filename>")
		os.Exit(0)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error reading file.")
		os.Exit(1)
	}

	dataString := string(data)

	for i, line := range strings.Split(dataString, "\n") {
		fmt.Println(i)

		for j, column := range strings.Split(line, ",") {
			fmt.Println("\t", j, "\t", column)
		}

	}


}