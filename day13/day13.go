package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solution(filename string) int {

	points, instructions := GetInput(filename)

	x := 0
	y := 0
	for _, point := range points {

		if point.X > x {
			x = point.X
		}
		if point.Y > y {
			y = point.Y
		}

	}
	grid := make([][]string, y+1)
	for i := range grid {
		grid[i] = make([]string, x+1)
	}

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			grid[i][j] = "."
		}

	}
	for _, point := range points {

		grid[point.Y][point.X] = "#"

	}

	instruction := instructions[0]
	if instruction.Edge == "x" {

		grid = FoldVertical(grid, instruction.Value)

	} else {

		grid = FoldHorizontal(grid, instruction.Value)
	}
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == "#" {
				counter++
			}
		}

	}

	return counter
}

func FoldVertical(grid [][]string, s string) [][]string {

	value, _ := strconv.Atoi(s)
	points := []Point{}
	for i := 0; i < len(grid); i++ {

		for j := value; j < len(grid[0]); j++ {

			if grid[i][j] == "#" {
				point := Point{X: j, Y: i}

				points = append(points, point)
			}
		}
	}

	newGrid := [][]string{}

	for _, v := range grid {

		newGrid = append(newGrid, v[:value])
	}

	for _, point := range points {

		grid[point.Y][len(grid[0])-point.X] = "#"

	}
	return newGrid
}

func FoldHorizontal(grid [][]string, s string) [][]string {

	value, _ := strconv.Atoi(s)
	points := []Point{}
	for i := value; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == "#" {
				point := Point{X: j, Y: i}

				points = append(points, point)
			}
		}
	}

	newGrid := grid[:value]

	for _, point := range points {

		grid[len(grid)-point.Y-1][point.X] = "#"

	}
	return newGrid

}

func SolutionPartTwo(filename string) int {

	points, instructions := GetInput(filename)

	x := 0
	y := 0
	for _, point := range points {

		if point.X > x {
			x = point.X
		}
		if point.Y > y {
			y = point.Y
		}

	}
	grid := make([][]string, y+1)
	for i := range grid {
		grid[i] = make([]string, x+1)
	}

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			grid[i][j] = "."
		}

	}
	for _, point := range points {

		grid[point.Y][point.X] = "#"

	}

	for _, instruction := range instructions {

		if instruction.Edge == "x" {

			grid = FoldVertical(grid, instruction.Value)

		} else {

			grid = FoldHorizontal(grid, instruction.Value)
		}
	}
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == "#" {
				counter++
			}
		}

	}

	return counter
}

func printGrid(grid [][]string) {

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == "#" {

				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println("")

	}

}

func GetInput(filename string) ([]Point, []Instruction) {

	input := []string{}
	points := []Point{}
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

	indexEmpty := 0
	for i, v := range input {

		if v == "" {

			indexEmpty = i
		}
	}

	stringNumbers := []string{}
	for i := 0; i < indexEmpty; i++ {

		stringNumbers = append(stringNumbers, strings.Split(input[i], ",")...)
	}

	numbers := []int{}
	numbers = StringArrayToInt(stringNumbers)

	for i := 0; i < len(numbers); i += 2 {

		point := Point{X: numbers[i], Y: numbers[i+1]}
		points = append(points, point)
	}
	instructionsString := []string{}

	for i := indexEmpty + 1; i < len(input); i++ {

		instructionsString = append(instructionsString, input[i])
	}

	instructions := []Instruction{}
	re := regexp.MustCompile("[0-9]+")
	for _, instruction := range instructionsString {

		instructionToAdd := Instruction{}
		if strings.ContainsAny(instruction, "x") {
			instructionToAdd.Edge = "x"
		} else {
			instructionToAdd.Edge = "y"
		}

		value := re.FindAllString(instruction, -1)

		instructionToAdd.Value = value[0]
		instructions = append(instructions, instructionToAdd)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return points, instructions
}

type Point struct {
	X int
	Y int
}

type Instruction struct {
	Value string
	Edge  string
}

func StringArrayToInt(parsedStr []string) []int {
	var intArray []int
	for _, v := range parsedStr {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("cannot parse str to int ")
		}
		intArray = append(intArray, number)
	}
	return intArray
}
