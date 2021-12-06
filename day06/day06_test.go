package day06_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day06"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution(80)

	expected := 5934

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo()

	expected := 26984457539

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
