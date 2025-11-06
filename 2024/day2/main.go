package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {
	if ok := utils.CheckOSArguments(); !ok {
		return
	}

	lines, err := utils.ReadLines(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	intArr, err := linesToIntArrays(lines)
	if err != nil {
		return
	}

	totalSafes := checkValidSafes(intArr)

	fmt.Printf("Total safes: %d\n", totalSafes)
}

func linesToIntArrays(lines []string) ([][]int, error) {
	var res [][]int

	for i := range lines {

		splitLines := strings.Split(lines[i], " ")
		curr := []int{}

		for _, num := range splitLines {
			val, err := strconv.Atoi(num)
			if err != nil {
				return [][]int{}, err
			}

			curr = append(curr, val)
		}

		res = append(res, curr)
	}

	return res, nil
}

func checkValidSafes(arr [][]int) int {
	var totalSafes int

	for _, list := range arr {

		ok := checkValidList(list)

		if !ok {
			continue
		}

		totalSafes++
	}

	return totalSafes
}

func checkValidList(arr []int) bool {
	for rightPtr := 1; rightPtr < len(arr); rightPtr++ {
		diff := int(math.Abs(float64(arr[rightPtr] - arr[rightPtr-1])))

		if diff >= 3 {
			return false
		}
	}
	return true
}
