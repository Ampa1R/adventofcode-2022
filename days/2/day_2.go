package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
	Rock | Paper | Scissors
	A    | B     | C
    X    | Y     | Z

	The score for a single round is the score for the shape you selected:
		1 for Rock
		2 for Paper
		3 for Scissors

	plus the score for the outcome of the round:
		0 if you lost
		3 if the round was a draw
		6 if you won


	1 = 1 = 0 // draw // (1 - 1) % 3 = 0
	1 < 2 = -1 // win // (1 - 2) % 3 = 2
	1 > 3 = -2 // lose // (1 - 3) % 3 = 1

	2 > 1 = 1 // lose // (2 - 1) % 3 = 1
	2 = 2 = 0 // draw // (2 - 2) % 3 = 0
	2 < 3 = -1 // win // (2 - 3) % 3 = 2

	3 < 1 = 2 // win // (3 - 1) % 3 = 2
	3 > 2 = 1 // lose // (3 - 2) % 3 = 1
	3 = 3 = 0 // draw // (3 - 3) % 3 = 0

	part 2
	lose | draw | win
	-----------------
	X    | Y    | Z

	A Y (draw) ->
		1 - 0 = 1
		1
		A X = 1 + 3 = 4
	B X (lose) ->
		2 - 1 = 1
		1
		B X = 1 + 0 = 1
	C Z (win) ->
		3 - 2 = 1
		1
		C X = 1 + 6 = 7

	A X (lose) ->
		1 - 1 = 0
		3
		A Z = 3 + 0 = 3

*/

func main() {
	solveFirstProblem()
	solveSecondProblem()
}

// Input data stuff
func loadInput() string {
	absPath, _ := filepath.Abs("days/2/input_2.txt")

	inputData, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	return string(inputData)
}

func parseInput(input string) (inputs [][]string) {
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		var rowInput = strings.Split(row, " ")
		inputs = append(inputs, rowInput)
	}

	return
}

func getInput() (input [][]string) {
	rawInput := loadInput()
	input = parseInput(rawInput)

	return
}

// scores part 1
func getScore(opponentPlay string, userPlay string) int {
	resultScore := 0

	scoresMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,

		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	x := scoresMap[opponentPlay]
	userPlayScore := scoresMap[userPlay]

	diff := (3 + (x - userPlayScore)) % 3

	switch diff {
	case 0: // draw
		resultScore = 3
	case 1: // lose
		resultScore = 0
	case 2: // win
		resultScore = 6
	}
	//fmt.Printf("Round result is %d + %d = %d (%s %s) \n", resultScore, userPlayScore, resultScore+userPlayScore, opponentPlay, userPlay)
	return resultScore + userPlayScore
}

func solveFirstProblem() {
	input := getInput()
	points := 0
	for _, round := range input {
		opponentPlay := round[0]
		userPlay := round[1]
		roundPoint := getScore(opponentPlay, userPlay)
		points += roundPoint
	}

	result := points
	fmt.Printf("First problem answer is %d \n", result)
	// 13924 is correct
}
func getRevertScore(opponentPlay string, result string) int {
	resultScore := 0

	scoresMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,

		"X": 1, // lose
		"Y": 0, // draw
		"Z": 2, // win
	}

	x := scoresMap[opponentPlay]
	resultWeight := scoresMap[result]
	userPlayScore := (3 + (x - resultWeight)) % 3

	if userPlayScore == 0 {
		userPlayScore = 3
	}

	switch result {
	case "X": // lose
		resultScore = 0
	case "Y": // draw
		resultScore = 3
	case "Z": // win
		resultScore = 6
	}
	//fmt.Printf("Round result is %d + %d = %d (%s %s) \n", resultScore, userPlayScore, resultScore+userPlayScore, opponentPlay, result)
	return resultScore + userPlayScore
}

func solveSecondProblem() {
	input := getInput()
	points := 0
	for _, round := range input {
		opponentPlay := round[0]
		userPlay := round[1]
		roundPoint := getRevertScore(opponentPlay, userPlay)
		points += roundPoint
	}

	result := points
	fmt.Printf("Second problem answer is %d \n", result)
	// 13448 is correct
}
