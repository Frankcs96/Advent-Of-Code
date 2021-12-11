package day11

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solution(filename string) int {
	input := GetInput(filename)
	cave := GetCave(input)
	solution := 0
	for i := 0; i < 100; i++ {

		explosions := []Position{}
		for i, row := range cave {

			for j := range row {

				if !isAExplosionPoint(&explosions, i, j) {
					cave[i][j] = cave[i][j] + 1
				}

				if cave[i][j] > 9 {
					Explosion(i, j, &cave, &explosions)
				}
			}
		}

		solution += len(explosions)
	}
	return solution
}
func SolutionPartTwo(filename string) int {

	input := GetInput(filename)
	cave := GetCave(input)

	isSync := false
	solution := 0

	for !isSync {

		explosions := []Position{}
		for i, row := range cave {

			for j := range row {

				if !isAExplosionPoint(&explosions, i, j) {
					cave[i][j] = cave[i][j] + 1
				}

				if cave[i][j] > 9 {
					Explosion(i, j, &cave, &explosions)
				}

			}
		}
		for i, row := range cave {

			for j := range row {

				if cave[i][j] != 0 {
					isSync = false
					break
				}

				isSync = true

			}
			if !isSync {
				break
			}
		}

		solution++
	}
	return solution

}

func Explosion(i, j int, cave *[10][10]int, explosions *[]Position) int {

	*explosions = append(*explosions, Position{i: i, j: j})

	cave[i][j] = 0
	if j-1 >= 0 {

		if !isAExplosionPoint(explosions, i, j-1) {
			cave[i][j-1] = cave[i][j-1] + 1
		}

		if cave[i][j-1] > 9 {
			Explosion(i, j-1, cave, explosions)
		}

	}
	if j+1 <= 9 {

		if !isAExplosionPoint(explosions, i, j+1) {
			cave[i][j+1] = cave[i][j+1] + 1
		}

		if cave[i][j+1] > 9 {
			Explosion(i, j+1, cave, explosions)
		}

	}
	if i+1 <= 9 {

		if !isAExplosionPoint(explosions, i+1, j) {
			cave[i+1][j] = cave[i+1][j] + 1
		}

		if cave[i+1][j] > 9 {
			Explosion(i+1, j, cave, explosions)
		}

	}
	if i-1 >= 0 {

		if !isAExplosionPoint(explosions, i-1, j) {
			cave[i-1][j] = cave[i-1][j] + 1
		}

		if cave[i-1][j] > 9 {
			Explosion(i-1, j, cave, explosions)
		}

	}
	if i-1 >= 0 && j-1 >= 0 {

		if !isAExplosionPoint(explosions, i-1, j-1) {
			cave[i-1][j-1] = cave[i-1][j-1] + 1
		}

		if cave[i-1][j-1] > 9 {
			Explosion(i-1, j-1, cave, explosions)
		}

	}
	if i-1 >= 0 && j+1 <= 9 {

		if !isAExplosionPoint(explosions, i-1, j+1) {
			cave[i-1][j+1] = cave[i-1][j+1] + 1
		}

		if cave[i-1][j+1] > 9 {
			Explosion(i-1, j+1, cave, explosions)
		}

	}
	if i+1 <= 9 && j-1 >= 0 {
		if !isAExplosionPoint(explosions, i+1, j-1) {
			cave[i+1][j-1] = cave[i+1][j-1] + 1
		}

		if cave[i+1][j-1] > 9 {

			Explosion(i+1, j-1, cave, explosions)
		}

	}
	if i+1 <= 9 && j+1 <= 9 {

		if !isAExplosionPoint(explosions, i+1, j+1) {
			cave[i+1][j+1] = cave[i+1][j+1] + 1
		}

		if cave[i+1][j+1] > 9 {
			Explosion(i+1, j+1, cave, explosions)
		}

	}
	return 2
}

func isAExplosionPoint(explosions *[]Position, i, j int) bool {

	isExplosion := false

	for _, point := range *explosions {

		if point.i == i && point.j == j {

			isExplosion = true
		}

	}
	return isExplosion
}

func GetCave(input []string) [10][10]int {

	var cave [10][10]int
	for i, line := range input {

		for j, val := range line {
			//int(val) return the ascii code so we substract 48 because we know that 0 is represented by 48
			number := int(val) - 48

			cave[i][j] = number
		}

	}
	return cave
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

type Position struct {
	i int
	j int
}
