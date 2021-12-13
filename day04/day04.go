package day04

import (
	"bufio"
	. "github.com/Frankcs96/Advent-Of-Code-2021-/utils"
	"log"
	"os"
	"regexp"
	"strings"
)

func Solution() int {
	input := ParseData("input.txt")

	numbers := GetNumbers(input)
	boards := GetBoards(input)

	validNumbersIndex := 5
	validNumbers := numbers[:validNumbersIndex]
	isWinner := false
	winnerBoardIndex := -1

	for !isWinner {

		isWinner, winnerBoardIndex = checkColumns(boards, validNumbers)
		isWinner, winnerBoardIndex = CheckRows(boards, validNumbers)
		if !isWinner {

			validNumbersIndex++
			validNumbers = numbers[:validNumbersIndex]
		}

	}

	sumUnmarkedNumbers := SumAllUnmarkedNumbers(boards[winnerBoardIndex], validNumbers)
	return sumUnmarkedNumbers * validNumbers[validNumbersIndex-1]
}

func SolutionPartTwo() int {
	input := ParseData("input.txt")

	numbers := GetNumbers(input)
	boards := GetBoards(input)

	validNumbersIndex := 5
	validNumbers := numbers[:validNumbersIndex]
	isColumnComplete := false
	isRowComplete := false
	isWinner := false
	winnerBoardIndex := -1

	for !isWinner {

		isColumnComplete, winnerBoardIndex = checkColumns(boards, validNumbers)
		isRowComplete, winnerBoardIndex = CheckRows(boards, validNumbers)

		if isColumnComplete || isRowComplete {
			isWinner = true
		}

		if isWinner {
			if len(boards) > 1 {
				isWinner = false
				var boardsLeft [][5][5]int
				for i, board := range boards {
					if i != winnerBoardIndex {
						boardsLeft = append(boardsLeft, board)
					}
					boards = nil
					boards = boardsLeft
				}
			}
		} else {
			validNumbersIndex++
			validNumbers = numbers[:validNumbersIndex]
		}

	}

	sumUnmarkedNumbers := SumAllUnmarkedNumbers(boards[0], validNumbers)
	return sumUnmarkedNumbers * validNumbers[validNumbersIndex-1]
}

func SumAllUnmarkedNumbers(board [5][5]int, validNumbers []int) int {

	sum := 0
	for _, i := range board {

		for _, j := range i {
			if !numberIsInArray(j, validNumbers) {
				sum += j
			}
		}
	}
	return sum
}

func numberIsInArray(number int, validNumbers []int) bool {

	for _, v := range validNumbers {
		if v == number {

			return true
		}

	}
	return false
}

func checkColumns(boards [][5][5]int, validNumbers []int) (bool, int) {
	var col []int

	for i := 0; i < len(boards); i++ {

		for k := 0; k < len(boards[i]); k++ {

			for j := 0; j < len(boards[i][k]); j++ {
				col = append(col, boards[i][j][k])
			}

			colIsComplete := CheckIfArrayContainsArray(col, validNumbers)
			if colIsComplete {
				return true, i
			}
			col = nil
		}
	}
	return false, -1
}

func CheckRows(boards [][5][5]int, validNumbers []int) (bool, int) {

	var row []int
	for i := 0; i < len(boards); i++ {

		for k := 0; k < len(boards[i]); k++ {

			for j := 0; j < len(boards[i][k]); j++ {
				row = append(row, boards[i][k][j])
			}

			rowIsComplete := CheckIfArrayContainsArray(row, validNumbers)
			if rowIsComplete {
				return true, i
			}
			row = nil
		}

	}
	return false, -1
}

func CheckIfArrayContainsArray(row []int, validNumbers []int) bool {

	isContained := true

	for _, rowVal := range row {
		if !isContained {
			break
		}
		for _, validNumber := range validNumbers {
			if rowVal == validNumber {
				isContained = true
				break
			} else {
				isContained = false
			}
		}
	}
	return isContained
}

func ParseData(filename string) []string {

	var input []string
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()
		input = append(input, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func GetNumbers(input []string) []int {
	dirtyNumbers := input[0]
	stringArray := strings.Split(dirtyNumbers, ",")

	numbers := StringArrayToInt(stringArray)
	return numbers
}

func GetBoards(input []string) [][5][5]int {
	boardInput := input[2:]
	var trimmedInput []string

	re := regexp.MustCompile("[0-9]+")

	for _, v := range boardInput {
		trimmedInput = append(trimmedInput, re.FindAllString(v, -1)...)

	}

	var board [5][5]int
	allNumbers := StringArrayToInt(trimmedInput)
	allNumbersIndex := 0
	numberOfBoards := len(allNumbers) / 25

	//3d array wth first time using it :_)
	var boards [][5][5]int
	for k := 0; k < numberOfBoards; k++ {

		for i := 0; i < len(board); i++ {

			for j := 0; j < len(board); j++ {

				board[i][j] = allNumbers[allNumbersIndex]
				allNumbersIndex++
			}

		}
		boards = append(boards, board)

	}

	return boards
}
