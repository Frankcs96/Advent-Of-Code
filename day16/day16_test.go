package day16_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/day16"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("example_input.txt")
	expected := 31

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 889

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
