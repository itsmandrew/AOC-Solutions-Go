package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FileToArray(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res []int

	for scanner.Scan() {
		curr := scanner.Text()

		currInt, err := strconv.Atoi(curr)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, currInt)
	}

	return res
}

func CountElementsGreater(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			result++
		}
	}

	return result
}

func CalculateWindowSums(nums []int) []int {
	var result []int

	for i := 2; i < len(nums); i++ {
		curr := nums[i] + nums[i-1] + nums[i-2]
		result = append(result, curr)
	}

	return result
}

func main() {
	arr := FileToArray("inputs/sample.txt")

	windowSums := CalculateWindowSums(arr)
	result := CountElementsGreater(windowSums)

	fmt.Println(result)
}
