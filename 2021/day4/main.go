package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileToArrays(path string) ([]int, [][][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	// Parsing drawn numbers
	scanner.Scan()
	numStrings := strings.Split(scanner.Text(), ",")

	for _, val := range numStrings {
		numInt, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, numInt)
	}

	var boards [][][]int
	var currentBoard [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if len(currentBoard) > 0 {
				boards = append(boards, currentBoard)
				currentBoard = [][]int{}
			}
			continue
		}

		fields := strings.Fields(line)
		var row []int

		for _, v := range fields {
			numInt, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, numInt)
		}
		currentBoard = append(currentBoard, row)
	}

	if len(currentBoard) > 0 {
		boards = append(boards, currentBoard)
	}

	return numbers, boards
}

func MarkBoard(board [][]int, target int) {
	for i, row := range board {
		for j, val := range row {
			if val == target {
				board[i][j] = -1
			}
		}
	}
}

func MarkAllBoardsAndGetWinnerAndWinningNumber(boards [][][]int, numsDrawn []int) (int, int) {
	for i := range numsDrawn {
		for j := range boards {
			MarkBoard(boards[j], numsDrawn[i])
			if CheckBingo(boards[j]) {
				return j, numsDrawn[i]
			}
		}
	}

	return -1, -1
}

func CheckBingo(board [][]int) bool {
	// Checking row
	for i := 0; i < 5; i++ {

		rowFlag := true
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				rowFlag = false
				break
			}
		}

		if rowFlag {
			return true
		}
	}

	// Checking columns
	for i := 0; i < 5; i++ {
		columnFlag := true
		for j := 0; j < 5; j++ {
			if board[j][i] != -1 {
				columnFlag = false
				break
			}
		}
		if columnFlag {
			return true
		}
	}

	return false
}

func CalculateScore(board [][]int, winningNumber int) int {
	runningSum := 0

	for _, row := range board {
		for _, val := range row {
			if val != -1 {
				runningSum += val
			}
		}
	}

	return winningNumber * runningSum
}

func RunThroughBoardsTillOneRemaining(boards [][][]int, numsDrawn []int) (int, int) {
	numberOfBoards := len(boards)
	alreadyWon := make(map[int]bool)

	boardCount := 0

	for i := range numberOfBoards {
		alreadyWon[i] = false
	}

	for i := range numsDrawn {
		for j := range boards {
			if alreadyWon[j] {
				continue
			}

			MarkBoard(boards[j], numsDrawn[i])
			if CheckBingo(boards[j]) {
				alreadyWon[j] = true
				boardCount++

				if boardCount == numberOfBoards {
					return j, numsDrawn[i]
				}
			}
		}
	}

	return -1, -1
}

func main() {
	numsDrawn, boards := ReadFileToArrays("inputs/test.txt")

	lastBoard, lastNumber := RunThroughBoardsTillOneRemaining(boards, numsDrawn)

	fmt.Printf("Winner: %d\n", lastBoard)
	for _, row := range boards[lastBoard] {
		fmt.Printf("%v\n", row)
	}

	score := CalculateScore(boards[lastBoard], lastNumber)

	fmt.Printf("Here's the score: %d\n", score)
}
