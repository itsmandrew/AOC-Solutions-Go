package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) != 1 {
		fmt.Printf("Expected 1 argument, got %d\n", len(os.Args[1:]))
		return
	}

}
