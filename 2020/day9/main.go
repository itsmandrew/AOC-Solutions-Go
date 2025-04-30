package main

import (
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

	hashMap := make(map[int]int) // Value : index

	for i := 0; i < 5; i++ {
		hashMap[intArr[i]] = i
	}

	initialIndex := 0

	for i := 5; i < len(arr); i++ {
		checkValid(hashMap, intArr[i])
		delete(hashMap, intArr[initialIndex])
		hashMap[intArr[i]] = i
	}

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
