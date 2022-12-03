package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	absPath, _ := filepath.Abs("days/1/input_1.txt")

	inputData, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	foodCalories := strings.Split(string(inputData), "\n")

	var elvesCals []int
	curElfCal := 0
	for _, s := range foodCalories {
		if s == "" {
			elvesCals = append(elvesCals, curElfCal)
			curElfCal = 0
		} else {
			i, _ := strconv.Atoi(s)
			curElfCal += i
		}
	}
	sort.Ints(elvesCals)
	fmt.Printf("Top calories: %d \n", elvesCals[len(elvesCals)-1])

	topThreeCals := elvesCals[len(elvesCals)-3:]
	topThreeCalsSum := 0
	for _, i := range topThreeCals {
		topThreeCalsSum += i
	}
	fmt.Printf("Top three calories sum: %d \n", topThreeCalsSum)
}
