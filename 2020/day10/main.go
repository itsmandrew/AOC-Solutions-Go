package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	ok := utils.CheckOSArguments()
	if !ok {
		os.Exit(1)
	}

	arr, err := utils.ReadLines(os.Args[1])

	if err != nil {
		os.Exit(1)
	}

	intArr, err := utils.StrArrayToInt(arr)

	if err != nil {
		os.Exit(1)
	}

	intArr = append(intArr, 0)

	slices.Sort(intArr)

	fmt.Printf("%v\n", intArr)

	numsSet := make(map[int]bool)

	for _, val := range intArr {
		numsSet[val] = true
	}

	diffMap := make(map[int][]int)
	length := len(intArr) - 1

	for i := 0; i < len(intArr); i++ {
		for j := 1; j < 4; j++ {
			key := intArr[i] + j
			if _, ok := numsSet[key]; ok {
				diffMap[j] = append(diffMap[j], key)
				break
			}
		}
	}

	diffMap[3] = append(diffMap[3], intArr[length]+3)

	fmt.Printf("%v\n", diffMap)

	for k := range diffMap {
		fmt.Printf("key %v : length -> %d\n", k, len(diffMap[k]))
	}
}
