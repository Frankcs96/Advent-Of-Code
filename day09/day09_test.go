package day09_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day09"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	expected := 15

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	expected := 1134

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 550

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 1100682

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestIsLowPoint(t *testing.T) {

	adyacents := []int{4, 3, 5}
	got := IsLowPoint(adyacents, 3)
	expected := false

	if got != expected {
		t.Errorf("fail expecting %v got %v", expected, got)
	}
	adyacents = []int{4, 3, 5, 6}
	got = IsLowPoint(adyacents, 2)
	expected = true

	if got != expected {
		t.Errorf("fail expecting %v got %v", expected, got)
	}
	adyacents = []int{4, 3, 5, 6}
	got = IsLowPoint(adyacents, 4)
	expected = false

	if got != expected {
		t.Errorf("fail expecting %v got %v", expected, got)
	}
	adyacents = []int{4, 1}
	got = IsLowPoint(adyacents, 1)
	expected = false

	if got != expected {
		t.Errorf("fail expecting %v got %v", expected, got)
	}

}
