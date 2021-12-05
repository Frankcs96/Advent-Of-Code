package main

import (
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution("input-example.txt")
	expected := 7

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input-example.txt")
	expected := 5

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolution(t *testing.T) {

	solution := Solution("input.txt")
	expected := 1298

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo("input.txt")
	expected := 1248

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
