package day07

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solution(filename string) int {

	input := GetInput(filename)

	m := map[int]int{}

	for _, v := range input {
		m[v] = 0
	}

	for k := range m {

		for _, v := range input {
			fuel := v - k
			m[k] += int(math.Abs(float64(fuel)))
		}
	}

	min := math.MaxInt
	for _, v := range m {

		if v < min {

			min = v
		}

	}

	return min
}

func SolutionPartTwo(filename string) int {

	input := GetInput(filename)

	m := map[int]int{}

	max, min := MaxMinArray(input)

	for i := min; i <= max; i++ {
		m[i] = 0
	}

	for k := range m {

		for _, v := range input {
			fuel := v - k

			absolute := math.Abs(float64(fuel))
			nNumberSum := SumNNumbers(int(absolute))
			m[k] += nNumberSum
		}
	}

	minOfMap := math.MaxInt
	for _, v := range m {

		if v < minOfMap {

			minOfMap = v
		}

	}

	return minOfMap
}

func SumNNumbers(number int) int {

	return number * (number + 1) / 2
}

func GetInput(filename string) []int {

	input := []int{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()

		for _, v := range strings.Split(val, ",") {

			number, err := strconv.Atoi(v)

			if err != nil {
				log.Fatal("could not parse str to int")
			}
			input = append(input, number)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
func MaxMinArray(array []int) (int, int) {

	max := array[0]
	min := array[0]

	for _, v := range array {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
