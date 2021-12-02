package main

import (
	"testing"
)

func TestExampleSolution(t *testing.T) {


  solution := Solution()
	expected := 150 

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionDayTwo(t *testing.T) {


  solution := SolutionPartTwo()
	expected := 900

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}


