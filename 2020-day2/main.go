package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		scanner := bufio.NewScanner(file)

		res := getSliceFromFile(scanner)

		// validPasswords := runPartOne(res)
		validPasswords := runPartTwo(res)

		file.Close()

		fmt.Printf("Valid passwords: %d\n", validPasswords)

	}
}

// func runPartOne(arr [][]string) int {
// 	validPasswords := 0

// 	for _, arr := range arr {

// 		lowest, highest := getFirstAndSecond(arr[0])
// 		want := getChar(arr[1])

// 		if check := matchChecker(want, lowest, highest, arr[2]); check {
// 			validPasswords += 1
// 		}

// 	}
// 	return validPasswords
// }

func runPartTwo(arr [][]string) int {
	validPasswords := 0

	for _, arr := range arr {
		first, second := getFirstAndSecond(arr[0])
		want := getChar(arr[1])

		if ok := matchChecker_v2(want, first, second, arr[2]); ok {
			validPasswords += 1
		}

	}
	return validPasswords

}

func getSliceFromFile(scan *bufio.Scanner) [][]string {

	res := [][]string{}

	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		res = append(res, line)
	}

	return res
}

func getFirstAndSecond(val string) (int, int) {

	arr := strings.Split(val, "-")

	lowest, err := strconv.Atoi(arr[0])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	highest, err := strconv.Atoi(arr[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	return lowest, highest
}

func getChar(val string) byte {

	res := val[0]

	return res
}

func matchChecker(want byte, lowest, highest int, val string) bool {

	wc := 0

	for _, char := range val {
		if byte(char) == want {
			wc += 1
		}
	}

	return (lowest <= wc) && (wc <= highest)
}

func matchChecker_v2(want byte, first, second int, val string) bool {

	cnt := 0
	if val[first-1] == want {
		cnt += 1
	}

	if val[second-1] == want {
		cnt += 1
	}

	if cnt == 1 {
		return true
	}
	return false
}
