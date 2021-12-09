package day08_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day08"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	// expected := 26
	expected := 26

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("example_input.txt")
	expected := 61229

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 247

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 933305

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestCalculateTopWire(t *testing.T) {

	var wires = []string{"abcf", "asdwer", "mblqop", "ab", "afvq", "abg"}

	got := CalculateTopWire(wires)
	expected := "g"

	if got != expected {
		t.Errorf("fail expecting %s got %s", expected, got)
	}

}
func TestCalculateBottomWire(t *testing.T) {

	var wires = []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}

	topWire := "d"
	got := CalculateBottomWire(wires, topWire)
	expected := "c"

	if got != expected {
		t.Errorf("fail expecting %s got %s", expected, got)
	}

}
