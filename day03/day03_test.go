package main

import (
	"testing"
)

func TestExampleSolution(t *testing.T) {


  solution := Solution()
  gamma := 22
  epsilon := 9
	expected := gamma * epsilon
  
  

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionDayTwo(t *testing.T) {


  solution := SolutionPartTwo()
	expected := 230

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

