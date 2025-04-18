package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments, exiting...")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	required := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
	}

	validCnt := 0

	m := map[string]string{}

	for scanner.Scan() {

		line := scanner.Text()

		if len(strings.TrimSpace(line)) > 0 {
			createMap(line, m)
		} else {
			fmt.Printf("the map: %v\n", m)

			if checkAllKeys(required, m) {
				validCnt++
			}
			m = map[string]string{}
		}
	}

	// fmt.Printf("the map: %v\n", m)

	if checkAllKeys(required, m) {
		validCnt++
	}

	fmt.Printf("valid passports: %d\n", validCnt)

}
func createMap(line string, m map[string]string) {
	arr := strings.Split(line, " ")

	for _, pair := range arr {
		two_pair := strings.Split(pair, ":")
		m[two_pair[0]] = two_pair[1]
	}
}

func checkAllKeys(req []string, m map[string]string) bool {

	for _, key := range req {
		if _, ok := m[key]; !ok {
			return false
		}
	}
	return true
}
