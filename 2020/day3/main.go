package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Printf("Got %d arguments, expected\n", len(os.Args[1:]))
		os.Exit(-1)
	}

}
