package utils

import (
	"log"
	"strconv"
)

func StringArrayToInt(stringArray []string) []int {

	var numbers []int
	for _, v := range stringArray {

		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("cannot cast string to int")
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func StringIsInArray(value string, array []string) bool {

	for _, v := range array {

		if v == value {

			return true
		}
	}
	return false
}
