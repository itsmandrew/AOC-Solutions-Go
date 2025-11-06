package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FileToArray(path string) []string {
	var result []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}

func CalculatePositionAndDepth(coords []string) (depth int, position int) {
	depth, position = 0, 0

	for _, str := range coords {
		parts := strings.Split(str, " ")

		numVal, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			position += numVal
		case "up":
			depth -= numVal
		case "down":
			depth += numVal
		}
	}

	return depth, position
}

func CalculatePositionAndDepthPart2(coords []string) (depth int, position int) {
	depth, position = 0, 0
	aim := 0

	for _, str := range coords {
		parts := strings.Split(str, " ")

		numVal, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "up":
			aim -= numVal

		case "down":
			aim += numVal

		case "forward":
			position += numVal
			depth += (aim * numVal)
		}
	}

	return depth, position
}

func main() {
	arr := FileToArray("inputs/test.txt")

	depth, position := CalculatePositionAndDepthPart2(arr)

	fmt.Println(depth * position)
}
