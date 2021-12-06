package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution() int {

	instructions := GetInstructions("input.txt")
	horizontal := 0
	depth := 0

	for _, v := range instructions {
		instruction := strings.Split(v, " ")

		command := instruction[0]
		unit, err := strconv.Atoi(instruction[1])

		if err != nil {
			log.Fatal("error parsing string to int")
		}

		if command == "forward" {

			horizontal += unit
		}

		if command == "down" {
			depth += unit
		}
		if command == "up" {
			depth -= unit
		}

	}
	solution := depth * horizontal
	return solution

}

func SolutionPartTwo() int {

	instructions := GetInstructions("input.txt")
	horizontal := 0
	depth := 0
	aim := 0

	for _, v := range instructions {
		instruction := strings.Split(v, " ")

		command := instruction[0]
		unit, err := strconv.Atoi(instruction[1])

		if err != nil {
			log.Fatal("error parsing string to int")
		}

		if command == "forward" {

			depth += aim * unit
			horizontal += unit
		}

		if command == "down" {
			aim += unit
		}
		if command == "up" {
			aim -= unit
		}

	}

	solution := depth * horizontal
	return solution

}

func GetInstructions(filename string) []string {

	report := []string{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()

		report = append(report, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return report
}
