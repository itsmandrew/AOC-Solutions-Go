package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToArray(path string) [][][]int {
	var result [][][]int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		splitLine := strings.Split(line, "->")
		var pairRes [][]int

		for _, pair := range splitLine {

			tuple := strings.Split(strings.TrimSpace(pair), ",")
			var curr []int

			for _, val := range tuple {

				valInt, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}

				curr = append(curr, valInt)
			}

			pairRes = append(pairRes, curr)
		}

		result = append(result, pairRes)

	}

	return result
}

func CreateGraphFromCoords(coords [][][]int) [][]int {
	maxX, maxY := 0, 0

	for _, pairs := range coords {
		for _, p := range pairs {
			if p[0] > maxX {
				maxX = p[0]
			}
			if p[1] > maxY {
				maxY = p[1]
			}
		}
	}

	graph := make([][]int, maxY+1)
	for i := range graph {
		graph[i] = make([]int, maxX+1)
	}

	return graph
}

func IsVertical(pairOne []int, pairTwo []int) bool {
	return pairOne[0] == pairTwo[0]
}

func IsHorizontal(pairOne []int, pairTwo []int) bool {
	return pairOne[1] == pairTwo[1]
}

func IsDiagonal(pairOne []int, pairTwo []int) bool {
	return AbsInt(pairOne[0]-pairTwo[0]) == AbsInt(pairOne[1]-pairTwo[1])
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MarkGraph(graph [][]int, coords [][][]int) {
	for _, pairs := range coords {
		if IsVertical(pairs[0], pairs[1]) {

			distance := AbsInt(pairs[0][1]-pairs[1][1]) + 1
			start := min(pairs[0][1], pairs[1][1])

			for i := 0; i < distance; i++ {
				graph[start+i][pairs[0][0]] += 1
			}

		} else if IsHorizontal(pairs[0], pairs[1]) {
			distance := AbsInt(pairs[0][0]-pairs[1][0]) + 1
			start := min(pairs[0][0], pairs[1][0])

			for i := 0; i < distance; i++ {
				graph[pairs[0][1]][start+i] += 1
			}

		} else if IsDiagonal(pairs[0], pairs[1]) {
			xStep := 1
			yStep := 1

			if pairs[0][0] > pairs[1][0] {
				xStep = -1
			}

			if pairs[0][1] > pairs[1][1] {
				yStep = -1
			}

			x, y := pairs[0][0], pairs[0][1]
			distance := AbsInt(pairs[0][0]-pairs[1][0]) + 1

			for i := 0; i < distance; i++ {
				graph[y][x] += 1
				x += xStep
				y += yStep
			}

		} else {
			continue
		}
	}
}

func CountGraph(graph [][]int, limit int) int {
	var result int

	for _, row := range graph {
		for _, cell := range row {
			if cell >= limit {
				result += 1
			}
		}
	}
	return result
}

func main() {
	arr := ReadFileToArray("inputs/test.txt")

	graph := CreateGraphFromCoords(arr)

	for _, pairs := range arr {
		fmt.Printf("%v\n", pairs)
	}

	MarkGraph(graph, arr)

	for _, row := range graph {
		fmt.Printf("%v\n", row)
	}

	result := CountGraph(graph, 2)
	fmt.Printf("Result: %d\n", result)
}
