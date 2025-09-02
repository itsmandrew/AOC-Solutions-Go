package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/sample.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int
	hashMap := make(map[int]int)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		values = append(values, num)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	res := 0

	for {
		for _, v := range values {
			res += v
			hashMap[res]++

			if hashMap[res] == 2 {
				fmt.Println("First repeated frequence:", res)
				return
			}
		}
	}
}
