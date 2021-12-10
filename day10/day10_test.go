package day10_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day10"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	expected := 26397

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	expected := 288957

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 436497

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 2377613374

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
