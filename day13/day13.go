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

	instruction := instructions[0]
	foldedPoints := make(map[Point]bool)
	if instruction.Edge == "x" {
		edgeLine, _ := strconv.Atoi(instruction.Value)
		for _, point := range points {
			if point.X > edgeLine {
				foldedPoint := Point{X: edgeLine*2 - point.X, Y: point.Y}
				foldedPoints[foldedPoint] = true
			} else {
				foldedPoints[point] = true
			}
		}
	} else {
		edgeLine, _ := strconv.Atoi(instruction.Value)
		for _, point := range points {
			if point.Y > edgeLine {
				foldedPoint := Point{X: point.X, Y: edgeLine*2 - point.Y}
				foldedPoints[foldedPoint] = true
			} else {
				foldedPoints[point] = true
			}
		}

	}
	return len(foldedPoints)
}

func SolutionPartTwo(filename string) int {
	points, instructions := GetInput(filename)

	var grid = make([][]int, 6)
	for i := range grid {
		grid[i] = make([]int, 39)
	}

	for _, point := range points {
		x := point.X
		y := point.Y
		for _, instruction := range instructions {
			if instruction.Edge == "x" {
				edgeLine, _ := strconv.Atoi(instruction.Value)

				if x > edgeLine {

					x = 2*edgeLine - x
				}
			} else {
				edgeLine, _ := strconv.Atoi(instruction.Value)
				if y > edgeLine {

					y = 2*edgeLine - y

				}

			}

		}
		grid[y][x] = 8
	}

	result := 0
	for _, row := range grid {
		for _, v := range row {
			if v == 8 {
				result++
				fmt.Print("█")
			} else {
				fmt.Print("░")

			}
		}

		fmt.Println("")
	}

	return result
}

func GetInput(filename string) ([]Point, []Instruction) {

	var input []string
	var points []Point
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

	var stringNumbers []string
	for i := 0; i < indexEmpty; i++ {

		stringNumbers = append(stringNumbers, strings.Split(input[i], ",")...)
	}

	var numbers []int
	numbers = StringArrayToInt(stringNumbers)

	for i := 0; i < len(numbers); i += 2 {

		point := Point{X: numbers[i], Y: numbers[i+1]}
		points = append(points, point)
	}
	var instructionsString []string

	for i := indexEmpty + 1; i < len(input); i++ {

		instructionsString = append(instructionsString, input[i])
	}

	var instructions []Instruction
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
