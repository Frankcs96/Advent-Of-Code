package main

import (
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
