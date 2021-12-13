package day05

import (
	"bufio"
	. "github.com/Frankcs96/Advent-Of-Code-2021-/utils"
	"log"
	"os"
	"regexp"
)

func Solution(isPartOne bool) int {

	filteredLines := GetInput("input.txt")

	if isPartOne {

		filteredLines = FilterDiagonalLines(filteredLines)
	}

	paths := make(map[Point]int)

	for _, line := range filteredLines {

		point := Point{}
		firstPoint := Point{X: line.X1, Y: line.Y1}
		paths[firstPoint] = paths[firstPoint] + 1

		for line.X1 != line.X2 || line.Y1 != line.Y2 {

			if line.X1 > line.X2 {
				// x1 goes down

				point.X = line.X1 - 1
				point.Y = line.Y1

				line.X1--

			}
			if line.Y1 > line.Y2 {
				point.X = line.X1
				point.Y = line.Y1 - 1
				line.Y1--
			}

			if line.X1 < line.X2 {
				point.X = line.X1 + 1
				point.Y = line.Y1
				line.X1++

			}
			if line.Y1 < line.Y2 {
				point.X = line.X1
				point.Y = line.Y1 + 1
				line.Y1++
			}

			paths[point] = paths[point] + 1

		}

	}

	solution := 0

	for _, v := range paths {

		if v > 1 {
			solution++
		}

	}
	return solution
}

func SolutionPartTwo() int {
	// My first solution works for part 2 too so just sending a flag to not filter diagonal lines
	return Solution(false)
}

func FilterDiagonalLines(input []Line) []Line {
	var filteredLines []Line

	for _, line := range input {

		if line.X1 == line.X2 || line.Y1 == line.Y2 {
			filteredLines = append(filteredLines, line)
		}
	}

	return filteredLines
}

func GetInput(filename string) []Line {

	var lines []Line
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()

		re := regexp.MustCompile("[0-9]+")
		parsedStr := re.FindAllString(val, -1)

		parsed := StringArrayToInt(parsedStr)
		for i := 0; i < len(parsed); i += 4 {

			line := Line{
				X1: parsed[i],
				Y1: parsed[i+1],
				X2: parsed[i+2],
				Y2: parsed[i+3],
			}
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

type Point struct {
	X int
	Y int
}
type Line struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}
