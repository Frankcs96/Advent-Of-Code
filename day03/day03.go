package day03

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Solution() int {

	report := GetReport("input.txt")
	gamma := ""
	epsilon := ""

	m := map[int]int{}
	// SOLUTION linear instead of quadratic
	for _, v := range report {

		for i, j := range v {

			if string(j) == "1" {
				m[i+1] = m[i+1] + 1
			}
		}
	}

	totalBits := len(report)
	var keys []int

	for k := range m {
		keys = append(keys, k)
	}
	//sort map keys
	sort.Ints(keys)

	for _, v := range keys {
		if m[v] > totalBits/2 {

			gamma += "1"
			epsilon += "0"

		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaDecimal) * int(epsilonDecimal)
}

func SolutionPartTwo() int {
	report := GetReport("input.txt")
	bitPosition := 0
	numberOfOnes := 0
	numberOfZeros := 0
	for len(report) > 1 {

		for _, number := range report {

			bit := number[bitPosition]

			if string(bit) == "1" {
				numberOfOnes++
			} else {
				numberOfZeros++
			}

		}

		var pv []string
		for _, number := range report {

			if numberOfZeros > numberOfOnes {
				if string(number[bitPosition]) == "0" {
					pv = append(pv, number)
				}
			} else {

				if string(number[bitPosition]) == "1" {
					pv = append(pv, number)
				}
			}
		}
		report = pv
		bitPosition++
		numberOfZeros = 0
		numberOfOnes = 0
		pv = nil
	}

	oxygen, _ := strconv.ParseInt(report[0], 2, 64)

	bitPosition = 0
	report = GetReport("input.txt")

	for len(report) > 1 {

		for _, number := range report {

			bit := number[bitPosition]

			if string(bit) == "1" {
				numberOfOnes++
			} else {
				numberOfZeros++
			}

		}

		var pv []string
		for _, number := range report {

			if numberOfZeros > numberOfOnes {
				if string(number[bitPosition]) == "1" {
					pv = append(pv, number)
				}
			} else {

				if string(number[bitPosition]) == "0" {
					pv = append(pv, number)
				}
			}
		}
		report = pv
		bitPosition++
		numberOfZeros = 0
		numberOfOnes = 0
		pv = nil
	}
	co2, _ := strconv.ParseInt(report[0], 2, 64)

	return int(co2 * oxygen)
}

func GetReport(filename string) []string {

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
