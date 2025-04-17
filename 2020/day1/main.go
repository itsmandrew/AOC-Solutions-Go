package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		scan := bufio.NewScanner(file)

		arr := getSliceFromFile(scan)

		partOneAnswer := bruteForce(arr)
		partTwoAnswer := bruteForce_v2(arr)

		fmt.Printf("Answer for part one: %d\n", partOneAnswer)
		fmt.Printf("Answer for part two: %d\n", partTwoAnswer)

		file.Close()
	}
}

func getSliceFromFile(scan *bufio.Scanner) []int {

	res := []int{}

	for scan.Scan() {
		s := scan.Text()

		val, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		res = append(res, val)
	}

	return res
}

func bruteForce(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j]
			}
		}
	}

	return -1
}

func bruteForce_v2(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}

	return -1
}
