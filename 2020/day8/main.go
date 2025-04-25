package main

import (
	"fmt"
	"os"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	if ok := utils.CheckOSArguments(); !ok {
		os.Exit(1)
	}

	data, err := utils.ReadLines(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%v\n", data)
}
