package day14_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day14"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	expected := 1588

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	var expected int64 = 2188189693529

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 2891

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	var expected int64 = 4607749009683

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
