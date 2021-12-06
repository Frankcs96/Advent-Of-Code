package day06

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution(days int) int {

	fishes := GetFishes("input.txt")
	m := map[int]int{}

	for _, fish := range fishes {

		m[fish] = m[fish] + 1

	}

	for i := 0; i < days; i++ {

		value8 := m[8]
		value7 := m[7]
		value6 := m[6]
		value5 := m[5]
		value4 := m[4]
		value3 := m[3]
		value2 := m[2]
		value1 := m[1]
		value0 := m[0]

		m[8] = value0
		m[7] = value8
		m[6] = value7 + value0
		m[5] = value6
		m[4] = value5
		m[3] = value4
		m[2] = value3
		m[1] = value2
		m[0] = value1
	}

	counter := 0
	for _, v := range m {

		counter += v
	}
	return counter
}

func SolutionPartTwo() int {

	return Solution(256)
}

func GetFishes(filename string) []int {

	var fishes []int
	var strfishes []string
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()

		strfishes = append(strfishes, strings.Split(val, ",")...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, strfish := range strfishes {

		fish, err := strconv.Atoi(strfish)

		if err != nil {
			log.Fatal("could not parse str to int")
		}

		fishes = append(fishes, fish)

	}

	return fishes
}
