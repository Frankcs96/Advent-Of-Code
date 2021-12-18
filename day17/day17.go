package day17

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"

	"github.com/Frankcs96/Advent-Of-Code-2021-/utils"
)

func Solution(filename string) int {

	input := GetInput(filename)

	re := regexp.MustCompile("-?[0-9]+")

	numbersString := re.FindAllString(input[0], -1)

	numbers := utils.StringArrayToInt(numbersString)

	targetArea := TargetArea{
		X1: numbers[0],
		X2: numbers[1],
		Y1: numbers[2],
		Y2: numbers[3],
	}

	dy := math.Abs(float64(targetArea.Y1)) - 1

	maxHeight := dy * (dy + 1) / 2

	return int(maxHeight)
}

func SolutionPartTwo(filename string) int {

	input := GetInput(filename)

	re := regexp.MustCompile("-?[0-9]+")

	numbersString := re.FindAllString(input[0], -1)

	numbers := utils.StringArrayToInt(numbersString)

	targetArea := TargetArea{
		X1: numbers[0],
		X2: numbers[1],
		Y1: numbers[2],
		Y2: numbers[3],
	}

	var yPossibleSpeeds []int

	xPossibleSpeeds := getPossibleXSpeeds(targetArea.X1, targetArea.X2)
	maxY := math.Abs(float64(targetArea.Y1)) - 1

	for i := targetArea.Y1; i <= int(maxY); i++ {
		yPossibleSpeeds = append(yPossibleSpeeds, i)
	}

	for i := targetArea.X1; i <= targetArea.X2; i++ {

		xPossibleSpeeds = append(xPossibleSpeeds, i)

	}

	var possibleSpeeds []Speed

	for _, speedX := range xPossibleSpeeds {

		for _, speedY := range yPossibleSpeeds {

			if isReachingTheTarget(speedX, speedY, targetArea) {

				possibleSpeeds = append(possibleSpeeds, Speed{X: speedX, Y: speedY})
			}
		}
	}

	return len(possibleSpeeds)
}

func isReachingTheTarget(speedX, speedY int, target TargetArea) bool {

	positionX := 0

	positionY := 0

	for positionX <= target.X2 && positionY >= target.Y1 {

		positionX += speedX
		positionY += speedY

		if speedX != 0 {

			speedX--
		}

		speedY--
		if positionX >= target.X1 && positionX <= target.X2 && positionY >= target.Y1 && positionY <= target.Y2 {

			return true

		}

	}

	return false

}

func getPossibleXSpeeds(initX, finalX int) []int {
	minXSpeed := getMinSpeedOfX(initX)
	var xSpeeds []int

	isPossible := true

	for isPossible {

		isPossible = false
		sum := 0
		for i := minXSpeed; i >= 0; i-- {

			sum += i

			if sum >= initX && sum <= finalX {
				isPossible = true
			}

		}
		sum = 0
		if isPossible {
			xSpeeds = append(xSpeeds, minXSpeed)
			minXSpeed++
		}
	}

	return xSpeeds

}
func getMinSpeedOfX(positionX int) int {
	speed := 0

	for speed*(speed+1)/2 < positionX {

		speed++
	}
	return speed
}

func GetInput(filename string) []string {

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

type TargetArea struct {
	X1, X2, Y1, Y2 int
}
type Speed struct {
	X, Y int
}
