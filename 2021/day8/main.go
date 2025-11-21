package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFileToArray(path string) [][]string {
	var res [][]string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		right := strings.TrimSpace(line[1])
		outputLine := strings.Split(right, " ")

		res = append(res, outputLine)
	}

	return res
}

func GetUniqueSegments(uniqueMap map[int]int, input [][]string) map[int]int {
	result := make(map[int]int)

	for _, arr := range input {
		for _, val := range arr {
			for k, v := range uniqueMap {
				if len(val) == v {
					result[k]++
				}
			}
		}
	}

	return result
}

func CountMap(countMap map[int]int) int {
	var res int
	for _, v := range countMap {
		res += v
	}

	return res
}

func main() {
	uniqueMap := map[int]int{
		1: 2, // Num : Length
		4: 4,
		7: 3,
		8: 7,
	}
	arr := ReadFileToArray("inputs/test.txt")

	resultMap := GetUniqueSegments(uniqueMap, arr)
	count := CountMap(resultMap)

	fmt.Printf("Total count: %d\n", count)
}
