package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main()  {
	
	if len(os.Args) != 2 {
		fmt.Println("USAGE fetch <url>")
		os.Exit(0)
	}

	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "GET error: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Body: \n%s\n", b)

}
