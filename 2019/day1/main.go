package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	if ok := utils.CheckOSArguments(); !ok {
		os.Exit(1)
	}

	data, err := utils.ReadLines(os.Args[1])

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	res := calculateFuelSum(data)
	fmt.Printf("Total fuel: %d\n", res)
}

func calculateFuelSum(arr []string) int64 {

	totalFuel := 0

	for _, val := range arr {
		num, err := strconv.Atoi(val)

		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		totalFuel += calculateRecursive(num)
	}

	return int64(totalFuel)
}

func calculateRecursive(val int) int {

	fuel := val/3 - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + calculateRecursive(fuel)
}
