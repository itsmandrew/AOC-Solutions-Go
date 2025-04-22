package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) == 0 || len(os.Args[1:]) > 1 {
		fmt.Printf("Want one arguments, got %d arguments\n", len(os.Args[1:]))
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	blocks := parseFile(data)
	res := 0
	for _, block := range blocks {
		res += getUniqueKeys(block)
	}

	fmt.Printf("sum of counts: %d\n", res)
}

func parseFile(data []byte) []string {

	blocks := strings.Split(string(data), "\n\n")

	return blocks

}

func getUniqueKeys(block string) int {
	allYes := 0

	arr := strings.Split(block, "\n")
	length := len(arr)

	hashMap := map[string]int{}

	for _, line := range arr {
		splitLine := strings.Split(line, "")
		for _, value := range splitLine {
			hashMap[value]++
		}
	}

	for key := range hashMap {
		if hashMap[key] == length {
			allYes++
		}
	}

	return allYes
}
