package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	if ok := utils.CheckOSArguments(); !ok {
		os.Exit(1)
	}

	arr, err := fileToArray(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%v\n", arr)

}

func fileToArray(path string) ([]int, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := []int{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		for _, val := range line {

			if num, err := strconv.Atoi(val); err != nil {
				continue
			} else {
				arr = append(arr, num)
			}
		}
	}

	return arr, scanner.Err()
}
