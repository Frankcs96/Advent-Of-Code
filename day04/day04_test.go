package day04_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day04"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution()

	expected := 4512

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionDayTwo(t *testing.T) {

	solution := SolutionPartTwo()
	expected := 1924

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestArrayContainArray(t *testing.T) {
	arr1 := []int{1, 4, 9, 16}
	arr2 := []int{2, 1, 4, 5, 22, 43, 4, 9, 16}

	expected := true
	result := CheckIfArrayContainsArray(arr1, arr2)

	if result != expected {
		t.Errorf("fail expecting %t got %t", expected, result)
	}

}
