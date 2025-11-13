package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToArray(path string) []int {
	var result []int
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		arr := strings.Split(line, ",")

		for _, val := range arr {

			valInt, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, valInt)
		}
	}
	return result
}

// O(days * state) -> Can prob cut this down to O(days)
func RunProgram(state []int, days int) int {
	for i := 0; i < days; i++ {
		for index := range state {
			if state[index] == 0 {
				state[index] = 6
				state = append(state, 8)

			} else {
				state[index] -= 1
			}
		}
	}

	return len(state)
}

func RunProgramOptimized(state []int, days int) int {
	result := 0
	hashMap := make(map[int]int)
	for _, val := range state {
		hashMap[val]++
	}
	for i := 0; i < days; i++ {

		newMap := make(map[int]int)

		for timer, count := range hashMap {
			if timer == 0 {
				newMap[6] += count
				newMap[8] += count
			} else {
				newMap[timer-1] += count
			}
		}
		hashMap = newMap
	}

	for key := range hashMap {
		result += hashMap[key]
	}

	return result
}

func main() {
	arr := ReadFileToArray("inputs/test.txt")
	fmt.Printf("%v\n", arr)

	result := RunProgramOptimized(arr, 256)

	fmt.Printf("Fish: %v\n", result)
}
