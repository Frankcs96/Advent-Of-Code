package day09

import (
	"bufio"
	. "github.com/Frankcs96/Advent-Of-Code-2021-/utils"
	"log"
	"os"
	"sort"
	"strconv"
)

func Solution(filename string) int {

	input := GetInput(filename)

	lowPoints := []LowPoint{}

	for i := 0; i < len(input); i++ {

		if i == 0 {

			lowPoints = append(lowPoints, LowPointsTopLine(input[i], input[i+1])...)

		}
		if i > 0 && i < len(input)-1 {

			lowPoints = append(lowPoints, CalculateLowPoints(i, input[i], input[i-1], input[i+1])...)
		}

		if i == len(input)-1 {

			lowPoints = append(lowPoints, LowPointsBotLine(i, input[i], input[i-1])...)
		}

	}

	solution := 0
	for _, v := range lowPoints {
		solution += v.Value + 1

	}
	return solution
}

func SolutionPartTwo(filename string) int {

	input := GetInput(filename)

	var lowPoints []LowPoint

	for i := 0; i < len(input); i++ {

		if i == 0 {

			lowPoints = append(lowPoints, LowPointsTopLine(input[i], input[i+1])...)

		}
		if i > 0 && i < len(input)-1 {

			lowPoints = append(lowPoints, CalculateLowPoints(i, input[i], input[i-1], input[i+1])...)
		}

		if i == len(input)-1 {

			lowPoints = append(lowPoints, LowPointsBotLine(i, input[i], input[i-1])...)
		}

	}
	var basinsLength []int
	for _, lowPoint := range lowPoints {

		var visitedPoints []VisitedPoint
		counter := 0
		for y := 0; y < len(input); y++ {

			for x := 0; x < len(input[y]); x++ {

				if lowPoint.PositionX == x && lowPoint.PositionY == y {

					counter = calculateBasins(x, y, input, lowPoint.Value, &visitedPoints)
					basinsLength = append(basinsLength, counter)
				}

			}

		}

	}

	sort.Ints(basinsLength)

	solution := 1
	for i := len(basinsLength) - 1; i > len(basinsLength)-4; i-- {

		solution *= basinsLength[i]
	}

	return solution
}

func calculateBasins(x, y int, input []string, value int, visitedPoints *[]VisitedPoint) int {

	leftPoint := "no more to the left"
	rightPoint := "no more to right"
	topPoint := "no more to top"
	botPoint := "no more to bot"

	if !isVisited(x, y, *visitedPoints) {

		visitedPoint := VisitedPoint{PointX: x, PointY: y}
		*visitedPoints = append(*visitedPoints, visitedPoint)

	}

	if x-1 >= 0 {

		leftPoint = string(input[y][x-1])
		intLeftPoint, _ := strconv.Atoi(leftPoint)
		if intLeftPoint > value && intLeftPoint != 9 {
			calculateBasins(x-1, y, input, intLeftPoint, visitedPoints)

		}
	}

	if x+1 < len(input[y]) {

		rightPoint = string(input[y][x+1])
		intRightPoint, _ := strconv.Atoi(rightPoint)
		if intRightPoint > value && intRightPoint != 9 {
			calculateBasins(x+1, y, input, intRightPoint, visitedPoints)
		}
	}
	if y-1 >= 0 {

		topPoint = string(input[y-1][x])
		intTopPoint, _ := strconv.Atoi(topPoint)
		if intTopPoint > value && intTopPoint != 9 {

			calculateBasins(x, y-1, input, intTopPoint, visitedPoints)
		}

	}
	if y+1 < len(input) {

		botPoint = string(input[y+1][x])
		intBotPoint, _ := strconv.Atoi(botPoint)
		if intBotPoint > value && intBotPoint != 9 {
			calculateBasins(x, y+1, input, intBotPoint, visitedPoints)
		}

	}
	return len(*visitedPoints)
}
func isVisited(x, y int, visitedPoints []VisitedPoint) bool {

	isVisited := false
	for _, point := range visitedPoints {

		if point.PointX == x && point.PointY == y {
			isVisited = true
		}
	}
	return isVisited

}
func LowPointsBotLine(y int, line, top string) []LowPoint {
	lowPoints := []LowPoint{}
	for i := 0; i < len(line); i++ {
		adyacents := []string{}
		if i == 0 {
			adyacents = append(adyacents, string(top[i]))
			adyacents = append(adyacents, string(line[i+1]))

		}

		if i > 0 && i < len(line)-1 {

			adyacents = append(adyacents, string(top[i]))
			adyacents = append(adyacents, string(line[i-1]))
			adyacents = append(adyacents, string(line[i+1]))

		}
		if i == len(line)-1 {

			adyacents = append(adyacents, string(line[i-1]))
			adyacents = append(adyacents, string(top[i]))

		}

		intAdyacents := StringArrayToInt(adyacents)
		point, err := strconv.Atoi(string(line[i]))
		if err != nil {
			log.Fatal("fail casting string to int ")
		}
		if IsLowPoint(intAdyacents, point) {

			lowPoint := LowPoint{
				Value:     point,
				PositionX: i,
				PositionY: y,
			}
			lowPoints = append(lowPoints, lowPoint)
		}
	}
	return lowPoints
}

func CalculateLowPoints(y int, line, top, bot string) []LowPoint {
	lowPoints := []LowPoint{}
	for i := 0; i < len(line); i++ {
		adyacents := []string{}
		if i == 0 {
			adyacents = append(adyacents, string(top[i]))
			adyacents = append(adyacents, string(bot[i]))
			adyacents = append(adyacents, string(line[i+1]))

		}

		if i > 0 && i < len(line)-1 {

			adyacents = append(adyacents, string(top[i]))
			adyacents = append(adyacents, string(bot[i]))
			adyacents = append(adyacents, string(line[i-1]))
			adyacents = append(adyacents, string(line[i+1]))

		}
		if i == len(line)-1 {

			adyacents = append(adyacents, string(line[i-1]))
			adyacents = append(adyacents, string(top[i]))
			adyacents = append(adyacents, string(bot[i]))

		}

		intAdyacents := StringArrayToInt(adyacents)
		point, err := strconv.Atoi(string(line[i]))
		if err != nil {
			log.Fatal("fail casting string to int ")
		}
		if IsLowPoint(intAdyacents, point) {
			lowPoint := LowPoint{
				Value:     point,
				PositionX: i,
				PositionY: y,
			}
			lowPoints = append(lowPoints, lowPoint)
		}
	}
	return lowPoints
}

func LowPointsTopLine(s1, s2 string) []LowPoint {

	lowPoints := []LowPoint{}
	for i := 0; i < len(s1); i++ {

		adyacents := []string{}
		if i == 0 {
			adyacents = append(adyacents, string(s1[i+1]))
			adyacents = append(adyacents, string(s2[i]))

		}

		if i > 0 && i < len(s1)-1 {

			adyacents = append(adyacents, string(s1[i+1]))
			adyacents = append(adyacents, string(s2[i]))
			adyacents = append(adyacents, string(s1[i-1]))

		}
		if i == len(s1)-1 {

			adyacents = append(adyacents, string(s2[i]))
			adyacents = append(adyacents, string(s1[i-1]))

		}

		intAdyacents := StringArrayToInt(adyacents)
		point, err := strconv.Atoi(string(s1[i]))
		if err != nil {
			log.Fatal("fail casting string to int ")
		}
		if IsLowPoint(intAdyacents, point) {

			lowPoint := LowPoint{
				Value:     point,
				PositionX: i,
				PositionY: 0,
			}
			lowPoints = append(lowPoints, lowPoint)
		}
	}
	return lowPoints
}

type LowPoint struct {
	Value     int
	PositionX int
	PositionY int
}

func GetInput(filename string) []string {

	input := []string{}
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

func IsLowPoint(adyacents []int, point int) bool {

	isLowPoint := true
	for _, adjacent := range adyacents {

		if point >= adjacent {
			isLowPoint = false
		}
	}
	return isLowPoint
}

type VisitedPoint struct {
	PointX int
	PointY int
}
