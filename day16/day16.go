package day16

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Solution(filename string) int {

	input := GetInput(filename)
	binary := HexToBinary(input)

	counter := 0
	Parse(binary, &counter)

	return counter

}

func Parse(binary string, counter *int) {

	if len(binary) < 8 {
		return
	}
	if binary == "" {
		return
	}
	version := BinaryToInt(binary[:3])
	*counter += version

	ID := BinaryToInt(binary[3:6])

	if ID == 4 {
		pointer := 6
		result := ""
		for binary[pointer]-48 == 1 {
			group := binary[pointer : pointer+5]
			result += group[1:5]
			pointer += 5
		}
		// result += lastGroup[1:5]
		pointer += 5
		newString := binary[pointer:]
		Parse(newString, counter)

	} else {

		typeId := binary[6:7]

		if typeId == "0" {
			lenghtOfSubPackets := BinaryToInt(binary[7 : 7+15])
			packets := binary[22 : 22+lenghtOfSubPackets]
			otherOne := binary[22+lenghtOfSubPackets:]
			Parse(packets, counter)
			Parse(otherOne, counter)

		}

		if typeId == "1" {
			//numberOfSubPackages := BinaryToInt(binary[7 : 7+11])
			newPackages := binary[7+11:]

			Parse(newPackages, counter)

		}
	}
}

func ParsePartTwo(binary string, counter *int, sum *int, mult *int, min *int, max *int, less *int, greater *int, equal *int) {

	if len(binary) < 8 {
		return
	}
	if binary == "" {
		return
	}
	version := BinaryToInt(binary[:3])
	*counter += version

	ID := BinaryToInt(binary[3:6])

	result := ""
	if ID == 4 {
		pointer := 6
		for binary[pointer]-48 == 1 {
			group := binary[pointer : pointer+5]
			result += group[1:5]
			pointer += 5
		}
		lastGroup := binary[pointer : pointer+5]
		result += lastGroup[1:5]
		pointer += 5
		newString := binary[pointer:]

		*sum += BinaryToInt(result)
		*mult *= BinaryToInt(result)
		if BinaryToInt(result) < *min {
			*min = BinaryToInt(result)
		}
		if BinaryToInt(result) > *max {
			*max = BinaryToInt(result)
		}
		if *less == -1 {
			*less = BinaryToInt(result)
		} else {
			if *less < BinaryToInt(result) {
				*less = 1
			} else {
				*less = 0
			}
		}
		if *greater == -1 {
			*greater = BinaryToInt(result)
		} else {
			if *greater > BinaryToInt(result) {
				*greater = 1
			} else {
				*greater = 0
			}
		}

		if *equal == -1 {
			*equal = BinaryToInt(result)
		} else {
			if *equal == BinaryToInt(result) {
				*equal = 1
			} else {
				*equal = 0
			}
		}
		ParsePartTwo(newString, counter, sum, mult, min, max, less, greater, equal)

	} else {

		typeId := binary[6:7]

		if typeId == "0" {
			lenghtOfSubPackets := BinaryToInt(binary[7 : 7+15])
			packets := binary[22 : 22+lenghtOfSubPackets]
			otherOne := binary[22+lenghtOfSubPackets:]

			ParsePartTwo(packets, counter, sum, mult, min, max, less, greater, equal)
			ParsePartTwo(otherOne, counter, sum, mult, min, max, less, greater, equal)

		}

		if typeId == "1" {
			// numberOfSubPackages := BinaryToInt(binary[7 : 7+11])
			newPackages := binary[7+11:]

			ParsePartTwo(newPackages, counter, sum, mult, min, max, less, greater, equal)

		}
	}

	if ID == 4 {
		fmt.Println(BinaryToInt(result))
	} else {

		if ID == 0 {
			fmt.Println("+")
		}
		if ID == 1 {
			fmt.Println("*")
		}
		if ID == 2 {
			fmt.Println("min")
		}
		if ID == 3 {
			fmt.Println("maximu")
		}
		if ID == 5 {
			fmt.Println(">")
		}
		if ID == 6 {
			fmt.Println("<")
		}
		if ID == 7 {
			fmt.Println("=")
		}
	}
}

func SolutionPartTwo(filename string) int {

	input := GetInput(filename)
	binary := HexToBinary(input)

	counter := 0
	sum := 0
	mult := 1
	min := math.MaxInt64
	max := 0
	less := -1
	greater := -1
	equal := -1
	ParsePartTwo(binary, &counter, &sum, &mult, &min, &max, &less, &greater, &equal)

	fmt.Println(equal)
	return 1
}

func BinaryToInt(binary string) int {

	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {

		log.Fatal("error parsing binary to decimal")
	}
	return int(decimal)
}
func HexToBinary(input []string) string {
	binary := ""

	m := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	for _, line := range input {

		for _, char := range line {

			binary += m[string(char)]
		}
	}
	return binary
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
