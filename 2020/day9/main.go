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

	hashMap := make(map[int]int) // Value : index

	for i := 0; i < 25; i++ {
		hashMap[intArr[i]] = i
	}

	initialIndex := 0
	lastVal := 0

	for i := 25; i < len(arr); i++ {
		if !checkValid(hashMap, intArr[i]) {
			lastVal = intArr[i]
			break
		}
		delete(hashMap, intArr[initialIndex])
		initialIndex++
		hashMap[intArr[i]] = i
	}

	fmt.Printf("%v\n", lastVal)

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
