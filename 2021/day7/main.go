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
		line := strings.Split(scanner.Text(), ",")

		for _, val := range line {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			result = append(result, valInt)
		}
	}

	return result
}

func main() {
	arr := ReadFileToArray("inputs/sample.txt")

	fmt.Printf("%v\n", arr)
}
