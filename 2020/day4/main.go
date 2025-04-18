package main

import (
	"fmt"
	"io"
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

	data, err := io.ReadAll(file)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	hashMap := parsePassports(string(data))

	fmt.Printf("%v\n", checkValid(hashMap, required))
}

func parsePassports(input string) []map[string]string {

	res := []map[string]string{}

	blocks := strings.Split(input, "\n\n")

	for _, block := range blocks {

		newBlock := strings.Fields(block)
		m := map[string]string{}

		for _, kv := range newBlock {
			keyPair := strings.Split(kv, ":")
			m[keyPair[0]] = keyPair[1]
		}
		res = append(res, m)
	}
	return res
}

func checkValid(hashMap []map[string]string, reqs []string) int {
	validPassports := 0

	for _, entry := range hashMap {
		okFlag := true
		fmt.Printf("%v\n", entry)
		for _, key := range reqs {
			if _, ok := entry[key]; !ok {
				okFlag = false
				break
			}
		}
		if okFlag {
			validPassports++
		}
	}

	return validPassports
}
