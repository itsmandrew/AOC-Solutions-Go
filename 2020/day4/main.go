package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Please give a file name, exiting...")
		return
	}

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

	// required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	hashMap := parsePassports(string(data))

	res := 0

	for _, entry := range hashMap {
		if checkValidFields(entry) {
			res++
		}
	}

	fmt.Printf("valid passports: %d\n", res)
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

func checkValidKeys(hashMap map[string]string, reqs []string) bool {

	var flag bool

	for _, key := range reqs {
		flag = true
		if _, ok := hashMap[key]; !ok {
			flag = false
			break
		}
	}
	return flag
}

func checkValidFields(hashMap map[string]string) bool {

	var validators = map[string]func(string) bool{

		"byr": func(v string) bool {
			if len(v) > 4 {
				return false
			}
			y, err := strconv.Atoi(v)
			return err == nil && 1920 <= y && y <= 2002
		},

		"iyr": func(v string) bool {
			if len(v) > 4 {
				return false
			}
			y, err := strconv.Atoi(v)
			return err == nil && 2010 <= y && y <= 2020
		},

		"eyr": func(v string) bool {
			if len(v) > 4 {
				return false
			}
			y, err := strconv.Atoi(v)
			return err == nil && 2020 <= y && y <= 2030
		},

		"hgt": func(v string) bool {
			if strings.HasSuffix(v, "cm") {
				x, err := strconv.Atoi(strings.TrimSuffix(v, "cm"))
				return err == nil && 150 <= x && x <= 193
			}

			if strings.HasSuffix(v, "in") {
				x, err := strconv.Atoi(strings.TrimSuffix(v, "in"))
				return err == nil && 59 <= x && x <= 76
			}
			return false
		},
		"hcl": func(v string) bool {
			m, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, v)
			return m
		},
		"ecl": func(v string) bool {
			switch v {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				return true
			}
			return false
		},
		"pid": func(v string) bool {
			m, _ := regexp.MatchString(`^[0-9]{9}$`, v)
			return m
		},
	}

	for key := range validators {
		val, ok := hashMap[key]

		if !ok {
			return false
		}

		check := validators[key]
		if !(check(val)) {
			return false
		}
	}

	return true
}
