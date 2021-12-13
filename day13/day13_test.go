package day13_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day13"
	"testing"
)

func TestExampleSolution(t *testing.T) {
	solution := Solution("example_input.txt")
	expected := 17

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 737

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 96

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	expected := 16

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
