package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func RunProgram(arr []int) int {
	nums := make(map[int]bool)

	for _, val := range arr {
		nums[val] = true
	}

	minFuel := math.Inf(1)
	for k := range nums {
		currFuel := 0
		for _, val := range arr {

			distance := AbsInt(k - val)
			currFuel += distance
		}

		minFuel = math.Min(minFuel, float64(currFuel))
	}

	return int(minFuel)
}

func CreateScaleMap(distance int) map[int]int {
	distMap := make(map[int]int)
	var running int

	for i := 1; i < distance+1; i++ {
		running += i
		distMap[i] = running
	}

	return distMap
}

func Triangular(n int) int {
	return n * (n + 1) / 2
}

func RunProgramPart2(arr []int) int {
	minVal := slices.Min(arr)
	maxVal := slices.Max(arr)

	minFuel := math.MaxInt

	for i := minVal; i < maxVal+1; i++ {
		currFuel := 0
		for _, val := range arr {
			distance := AbsInt(val - i)
			currFuel += Triangular(distance)
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func main() {
	arr := ReadFileToArray("inputs/test.txt")

	fmt.Printf("%v\n", arr)

	minFuel := RunProgram(arr)
	minFuel2 := RunProgramPart2(arr)

	fmt.Printf("Minimum fuel: %d\n", minFuel)
	fmt.Printf("Part 2: %d\n", minFuel2)
}
