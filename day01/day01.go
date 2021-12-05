package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Solution(filename string) int {
	solution := 0
	report := GetReport(filename)
	for i := 1; i < len(report); i++ {

		if report[i] > report[i-1] {
			solution++

		}
	}
	return solution
}

func SolutionPartTwo(filename string) int {

	solution := 0
	report := GetReport(filename)

	for i := 0; i < len(report)-3; i++ {
		group1 := report[i] + report[i+1] + report[i+2]
		group2 := report[i+1] + report[i+2] + report[i+3]

		if group2 > group1 {
			solution++
		}

	}

	return solution
}

func GetReport(filename string) []int {

	report := []int{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal("error parsing string to int")
		}

		report = append(report, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return report
}
