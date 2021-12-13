package day05_test

import (
	"reflect"
	"testing"

	. "github.com/Frankcs96/Advent-Of-Code-2021-/day05"
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

func TestFilterHorizontalLines(t *testing.T) {

	lines := []Line{{X1: 20, Y1: 31, X2: 11, Y2: 41}, {X1: 98, Y1: 11, X2: 98, Y2: 41}, {X1: 3, Y1: 1234, X2: 0, Y2: 998}, {X1: 1234, Y1: 51, X2: 915, Y2: 51}}

	filterLines := FilterDiagonalLines(lines)

	expected := []Line{{X1: 98, Y1: 11, X2: 98, Y2: 41}, {X1: 1234, Y1: 51, X2: 915, Y2: 51}}

	if !reflect.DeepEqual(expected, filterLines) {
		t.Errorf("fail expecting %d got %d", expected, filterLines)
	}

}
