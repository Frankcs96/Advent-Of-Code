package day07_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day07"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("input-example.txt")
	expected := 37

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input-example.txt")
	expected := 168

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 340056

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 96592275

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestMinMaxArray(t *testing.T) {

	array := []int{34, 23, 50, 42, 42, 33, 33, 23, 50}
	gotMax, gotMin := MaxMinArray(array)

	expectedMax := 50
	expectedMin := 23

	if expectedMax != gotMax {
		t.Errorf("fail expecting %d got %d", expectedMax, gotMax)
	}
	if expectedMin != gotMin {
		t.Errorf("fail expecting %d got %d", expectedMin, gotMin)
	}

}
func TestSumNNumbers(t *testing.T) {

	expected := 21
	got := SumNNumbers(6)

	if got != expected {
		t.Errorf("fail expecting %d got %d", expected, got)
	}
	got = SumNNumbers(23)
	expected = 276
	if got != expected {
		t.Errorf("fail expecting %d got %d", expected, got)
	}

	got = SumNNumbers(9)
	expected = 45
	if got != expected {
		t.Errorf("fail expecting %d got %d", expected, got)
	}
}
