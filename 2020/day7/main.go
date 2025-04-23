package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

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

	result := countValidBags(graph)
	fmt.Printf("valid bags: %d\n", result)

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

func extractColors(line string) []string {
	words := strings.Fields(line)
	colors := []string{}

	for i := 0; i < len(words); {
		if words[i] == "no" {
			break
		}
		if i+2 < len(words) {
			color := words[i+1] + " " + words[i+2]
			colors = append(colors, color)
			i += 3
		} else {
			break
		}
	}
	return colors
}

func createMapping(data [][]string) map[string][]string {

	hashMap := make(map[string][]string)

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
