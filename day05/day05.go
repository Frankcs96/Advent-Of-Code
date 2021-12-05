package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Solution(isPartOne bool) int {

	filteredLines := GetInput("input.txt")

	if isPartOne {

		filteredLines = FilterDiagonalLines(filteredLines)
	}

	paths := make(map[Point]int)

	for _, line := range filteredLines {

		point := Point{}
		firstPoint := Point{x: line.x1, y: line.y1}
		paths[firstPoint] = paths[firstPoint] + 1

		for line.x1 != line.x2 || line.y1 != line.y2 {

			if line.x1 > line.x2 {
				// x1 goes down

				point.x = line.x1 - 1
				point.y = line.y1

				line.x1--

			}
			if line.y1 > line.y2 {
				point.x = line.x1
				point.y = line.y1 - 1
				line.y1--
			}

			if line.x1 < line.x2 {
				point.x = line.x1 + 1
				point.y = line.y1
				line.x1++

			}
			if line.y1 < line.y2 {
				point.x = line.x1
				point.y = line.y1 + 1
				line.y1++
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

		if line.x1 == line.x2 || line.y1 == line.y2 {
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

		parsed := stringArrayToInt(parsedStr)
		for i := 0; i < len(parsed); i += 4 {

			line := Line{
				x1: parsed[i],
				y1: parsed[i+1],
				x2: parsed[i+2],
				y2: parsed[i+3],
			}
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func stringArrayToInt(parsedStr []string) []int {
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

type Point struct {
	x int
	y int
}
type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}
