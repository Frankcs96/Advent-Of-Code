package main

import (
	"reflect"
	"testing"
)

func TestExampleSolution(t *testing.T) {

	solution := Solution(true)

	expected := 5

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}
func TestExampleSolutionPartTwo(t *testing.T) {

	solution := SolutionPartTwo()

	expected := 12

	if solution != expected {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestStringArrayToInt(t *testing.T) {

	input := []string{"2", "3", "123", "1394"}
	expected := []int{2, 3, 123, 1394}

	solution := stringArrayToInt(input)

	if !reflect.DeepEqual(expected, solution) {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestFilterHorizontalLines(t *testing.T) {

	lines := []Line{{x1: 20, y1: 31, x2: 11, y2: 41}, {x1: 98, y1: 11, x2: 98, y2: 41}, {x1: 3, y1: 1234, x2: 0, y2: 998}, {x1: 1234, y1: 51, x2: 915, y2: 51}}

	filterLines := FilterDiagonalLines(lines)

	expected := []Line{{x1: 98, y1: 11, x2: 98, y2: 41}, {x1: 1234, y1: 51, x2: 915, y2: 51}}

	if !reflect.DeepEqual(expected, filterLines) {
		t.Errorf("fail expecting %d got %d", expected, filterLines)
	}

}
