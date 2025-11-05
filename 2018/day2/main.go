package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("inputs/final.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	two, three := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		hashMap := make(map[rune]int)

		for _, c := range line {
			hashMap[c]++
		}

		twoFlag, threeFlag := false, false

		for _, v := range hashMap {
			if v == 2 {
				twoFlag = true
			}

			if v == 3 {
				threeFlag = true
			}
		}

		if twoFlag {
			two++
		}

		if threeFlag {
			three++
		}
	}

	checkSum := two * three

	fmt.Printf("two: %d, three: %d\n", two, three)
	fmt.Printf("Checksum: %d\n", checkSum)
}
