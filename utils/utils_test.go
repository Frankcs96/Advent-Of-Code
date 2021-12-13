package utils_test

import (
	. "github.com/Frankcs96/Advent-Of-Code-2021-/utils"
	"reflect"
	"testing"
)

func TestStringArrayToInt(t *testing.T) {

	input := []string{"2", "3", "123", "1394"}
	expected := []int{2, 3, 123, 1394}

	solution := StringArrayToInt(input)

	if !reflect.DeepEqual(expected, solution) {
		t.Errorf("fail expecting %d got %d", expected, solution)
	}

}

func TestStringIsInArray(t *testing.T) {

	input := []string{"dog", "cat", "salamander", "owl"}

	got := StringIsInArray("dog", input)

	if got != true {
		t.Errorf("fail expecting %t got %t", true, got)
	}

}
