package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

type Tuple struct {
	Command string
	Value   int
}

func main() {

	if ok := utils.CheckOSArguments(); !ok {
		os.Exit(1)
	}

	data, err := utils.ReadLines(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	arr := createMapping(data)

	res := getAccumulator(arr)
	fmt.Printf("%v\n", res)
}

func createMapping(lines []string) []Tuple {

	var arr []Tuple

	for _, line := range lines {
		newLine := strings.Split(line, " ")

		newNum, err := parseStringToInt(newLine[1])

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		tuple := Tuple{newLine[0], newNum}

		arr = append(arr, tuple)
	}

	return arr
}

func parseStringToInt(val string) (int, error) {
	return strconv.Atoi(val)
}

func getAccumulator(arr []Tuple) int {

	visited := map[int]bool{}

	accumulator := 0
	i := 0

	for {

		if _, ok := visited[i]; ok {
			break
		}
		visited[i] = true

		switch arr[i].Command {
		case "nop":
			i++
		case "acc":
			accumulator += arr[i].Value
			i++
		case "jmp":
			i += arr[i].Value
		}

	}

	return accumulator

}
