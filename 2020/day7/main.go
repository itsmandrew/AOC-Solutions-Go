package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

type BagRule struct {
	Name  string
	Count int
}

func main() {

	if ok := utils.CheckOSArguments(); !ok {
		os.Exit(1)
	}

	data, err := utils.ReadLines(os.Args[1])

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	cleanData := splitCleanAndData(data)

	graph := createMapping(cleanData)

	fmt.Println(countBagsInside("shiny gold", graph))

}

func splitCleanAndData(data []string) [][]string {
	splitLines := [][]string{}

	toRemove := []string{"bags", "bag", ".", ","}
	cleanLines := []string{}

	for _, val := range data {
		cleanLine := val
		for _, delim := range toRemove {
			cleanLine = strings.ReplaceAll(cleanLine, delim, "")
			cleanLine = strings.Trim(cleanLine, " ")
		}
		cleanLines = append(cleanLines, cleanLine)
	}

	for _, line := range cleanLines {
		splitLine := strings.Split(line, "contain")
		splitLines = append(splitLines, splitLine)
	}

	return splitLines
}

func countBagsInside(bag string, rules map[string][]BagRule) int {
	total := 0

	for _, rule := range rules[bag] {
		total += rule.Count * (1 + countBagsInside(rule.Name, rules))
	}

	return total
}

func extractColors(line string) []BagRule {
	words := strings.Fields(line)
	rules := []BagRule{}

	for i := 0; i < len(words); {
		if words[i] == "no" {
			break
		}
		if i+2 < len(words) {
			count, err := strconv.Atoi(words[i])
			if err != nil {
				i++
				continue
			}
			color := words[i+1] + " " + words[i+2]
			rules = append(rules, BagRule{Name: color, Count: count})
			i += 3

		} else {
			break
		}
	}
	return rules
}

func createMapping(data [][]string) map[string][]BagRule {

	hashMap := make(map[string][]BagRule)

	for _, arr := range data {
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		hashMap[key] = extractColors(value)
	}
	return hashMap
}

func depthFirstSearch(graph map[string][]string, node string, visited map[string]bool) bool {
	if node == "shiny gold" {
		return true
	}
	if visited[node] {
		return false
	}
	visited[node] = true

	for _, neighbor := range graph[node] {
		if depthFirstSearch(graph, neighbor, visited) {
			return true
		}
	}
	return false
}

func countValidBags(graph map[string][]string) int {
	count := 0
	for node := range graph {
		if node == "shiny gold" {
			continue
		}
		visited := make(map[string]bool)
		if depthFirstSearch(graph, node, visited) {
			count++
		}
	}
	return count
}
