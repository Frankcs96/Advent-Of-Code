package day03_test

import (
	"testing"

	. "github.com/Frankcs96/Advent-Of-Code-2021-/day03"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution()
	gamma := 22
	epsilon := 9
	expected := gamma * epsilon

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionDayTwo(t *testing.T) {

	solution := SolutionPartTwo()
	expected := 230

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
