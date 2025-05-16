package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/itsmandrew/aoc-go/utils"
)

func main() {

	ok := utils.CheckOSArguments()

	if !ok {
		return
	}

	lines, err := utils.ReadLines(os.Args[1])

	if err != nil {
		return
	}

	var cleanArr [][]string
	for _, val := range lines {
		splitLine := strings.Split(val, "   ")

		cleanArr = append(cleanArr, splitLine)
	}

	intArr, err := flattenArray(cleanArr)

	if err != nil {
		fmt.Println(err)
		return
	}

	arr1, arr2, err := getTwoArrays(intArr)

	if err != nil {
		fmt.Println(err)
		return
	}

	totalDistance, err := getTotalDistance(arr1, arr2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Total distance: %d\n", totalDistance)

}

func flattenArray(arr [][]string) ([]int, error) {

	var res []int

	for _, v := range arr {
		for _, n := range v {

			num, err := strconv.Atoi(n)
			if err != nil {
				return []int{}, err
			}
			res = append(res, num)
		}
	}

	return res, nil
}

func getTwoArrays(arr []int) ([]int, []int, error) {

	if len(arr)%2 != 0 {
		return []int{}, []int{}, fmt.Errorf("cannot perform split on uneven array")
	}

	var res1 []int
	var res2 []int

	for i := range arr {
		if i%2 == 0 {
			res1 = append(res1, arr[i])
		} else {
			res2 = append(res2, arr[i])
		}
	}

	return res1, res2, nil

}

func getTotalDistance(arr1, arr2 []int) (int, error) {

	if len(arr1) != len(arr2) {
		return 0, fmt.Errorf("Array1 and array2 are not of even length")
	}

	var totalDistance int

	slices.Sort(arr1)
	slices.Sort(arr2)

	for i := range arr1 {

		distance := int(math.Abs(float64(arr1[i] - arr2[i])))

		totalDistance += distance

	}

	return totalDistance, nil
}
