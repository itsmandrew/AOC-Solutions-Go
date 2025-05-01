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

	arr, err := utils.ReadLines(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	intArr, err := utils.StrArrayToInt(arr)

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	result := slidingWindow(intArr, 25)
	fmt.Printf("last value: %d\n", result)

}

func slidingWindow(arr []int, windowSize int) int {

	window := make(map[int]int)

	for i := 0; i < windowSize; i++ {
		window[arr[i]] = i
	}

	initialIndex := 0
	result := 0

	for i := windowSize; i < len(arr); i++ {
		if !checkValid(window, arr[i]) {
			result = arr[i]
			break
		}
		delete(window, arr[initialIndex])
		initialIndex++
		window[arr[i]] = i
	}

	return result

}

func checkValid(hashMap map[int]int, curr int) bool {

	for k, v := range hashMap {
		diff, ok := hashMap[curr-k]

		if ok && hashMap[diff] != v {
			return true
		}
	}

	return false
}
