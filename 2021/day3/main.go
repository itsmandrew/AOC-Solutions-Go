package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadFileToMap(path string) (map[int]int, int) {
	// Counts how many 0 bits there are per column
	hashMap := make(map[int]int)
	var lines int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for index, char := range line {
			if char == '0' {
				hashMap[index]++
			}
		}
		lines++
	}

	return hashMap, lines
}

func ReadFileToArray(path string) []string {
	var result []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func CalculateGammaAndEpsilon(bitMap map[int]int, lines int) (gammaRate string, epsilonRate string) {
	ones, zeroes := 0, 0

	for i := 0; i < len(bitMap); i++ {
		zeroes = bitMap[i]
		ones = lines - bitMap[i]

		if zeroes > ones {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	return gammaRate, epsilonRate
}

func CalculatePower(gammaRate string, epsilonRate string) int {
	gammaDecimal, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	epsilonDecimal, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(gammaDecimal) * int(epsilonDecimal)
}

func FilterSlice(bitArr []string, index int, key rune) []string {
	var newArr []string

	for _, val := range bitArr {
		if rune(val[index]) == key {
			newArr = append(newArr, val)
		}
	}
	return newArr
}

func FindMajorityKey(bitArr []string, index int) rune {
	var ones, zeros int

	for _, val := range bitArr {
		if val[index] == '1' {
			ones++
		} else {
			zeros++
		}
	}

	if ones >= zeros {
		return '1'
	}
	return '0'
}

func FindMinorityKey(bitArr []string, index int) rune {
	var ones, zeros int

	for _, val := range bitArr {
		if val[index] == '1' {
			ones++
		} else {
			zeros++
		}
	}

	if zeros <= ones {
		return '0'
	}
	return '1'
}

func RunPartTwoProgram(bitArr []string) int {
	o2GenArr, CO2ScrubberArr := bitArr, bitArr

	// For O2 Generator Rate
	for i := 0; len(o2GenArr) > 1; i++ {
		key := FindMajorityKey(o2GenArr, i)
		o2GenArr = FilterSlice(o2GenArr, i, key)
	}

	o2GenRate := o2GenArr[0]

	// For CO2 Scrubber Rate
	for i := 0; len(CO2ScrubberArr) > 1; i++ {
		key := FindMinorityKey(CO2ScrubberArr, i)
		CO2ScrubberArr = FilterSlice(CO2ScrubberArr, i, key)
	}

	co2ScrubberRate := CO2ScrubberArr[0]

	lifeSupportRating := CalculatePower(o2GenRate, co2ScrubberRate)

	return lifeSupportRating
}

func main() {
	hashMap, lines := ReadFileToMap("inputs/test.txt")
	bitArray := ReadFileToArray("inputs/test.txt")

	fmt.Printf("Total lines: %d\n", lines)

	gammaRate, epsilonRate := CalculateGammaAndEpsilon(hashMap, lines)
	power := CalculatePower(gammaRate, epsilonRate)

	fmt.Printf("Power: %d\n", power)

	oxygenGenRate := RunPartTwoProgram(bitArray)
	fmt.Printf("life support rating: %d\n", oxygenGenRate)
}
