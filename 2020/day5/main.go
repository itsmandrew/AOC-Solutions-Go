package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {

	if len(os.Args[1:]) != 1 {
		fmt.Printf("Expected 1 arguments, got %d\n arguments", len(os.Args[1:]))
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer file.Close()
	seatIDS := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		firstHalf, secondHalf := splitLine(line)
		row, col := binaryPartition(0, 127, firstHalf), binaryPartition(0, 7, secondHalf)
		seatIDS = append(seatIDS, (row*8)+col)
	}

	max := slices.Max(seatIDS)
	fmt.Printf("Max seat id: %d\n", max)

	mySeatID := findSeatID(seatIDS)
	fmt.Printf("My seat id: %d\n", mySeatID)
}

func splitLine(line string) ([]rune, []rune) {
	return []rune(line[:7]), []rune(line[7:])
}

func binaryPartition(low, high int, arr []rune) int {

	mid := 0

	for _, v := range arr {
		mid = (low + high) / 2
		if v == 'F' || v == 'L' {
			high = mid
		} else {
			low = mid + 1
		}

	}
	return low
}

func findSeatID(arr []int) int {
	slices.Sort(arr)

	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1]+1 != arr[i+1]-1 {
			return arr[i] + 1
		}
	}
	return -1
}
