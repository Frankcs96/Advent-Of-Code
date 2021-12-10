package day10

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func Solution(filename string) int {

	input := GetInput(filename)

	var stack Stack

	solution := 0
	for _, line := range input {

		for _, rune := range line {

			symbol := string(rune)

			if IsCloseChar(symbol) {
				lastChar := stack[len(stack)-1]

				if lastChar != "{" && symbol == "}" {

					solution += 1197
				}
				if lastChar != "<" && symbol == ">" {

					solution += 25137
				}
				if lastChar != "[" && symbol == "]" {

					solution += 57
				}
				if lastChar != "(" && symbol == ")" {

					solution += 3
				}
				stack.Pop()
			} else {

				stack.Push(symbol)
			}
		}
	}

	return solution
}

func SolutionPartTwo(filename string) int {

	input := GetInput(filename)
	scores := []int{}
	for _, line := range input {

		solution := 0
		var stack Stack
		isCorrupted := false
		for _, rune := range line {

			symbol := string(rune)

			if IsCloseChar(symbol) {
				lastChar := stack[len(stack)-1]

				if lastChar != "{" && symbol == "}" {

					isCorrupted = true
				}
				if lastChar != "<" && symbol == ">" {

					isCorrupted = true
				}
				if lastChar != "[" && symbol == "]" {

					isCorrupted = true
				}
				if lastChar != "(" && symbol == ")" {

					isCorrupted = true
				}
				stack.Pop()
			} else {

				stack.Push(symbol)
			}
		}
		if !isCorrupted {

			for !stack.IsEmpty() {
				lastValue := stack[len(stack)-1]

				solution = solution * 5
				if lastValue == "{" {
					solution += 3
				}
				if lastValue == "(" {
					solution += 1
				}
				if lastValue == "<" {
					solution += 4
				}
				if lastValue == "[" {
					solution += 2
				}
				stack.Pop()
			}

			scores = append(scores, solution)

		}

	}
	sort.Ints(scores)

	return scores[len(scores)/2]
}

func IsCloseChar(char string) bool {

	if char == "}" || char == ">" || char == "]" || char == ")" {

		return true
	}

	return false

}

func GetInput(filename string) []string {

	input := []string{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()
		input = append(input, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

// Do not use this stack implementation in a serious project :P
type Stack []string

func (s *Stack) Push(str string) {

	*s = append(*s, str)

}

func (s *Stack) IsEmpty() bool {

	return len(*s) == 0

}
func (s *Stack) Pop() {

	n := len(*s)

	*s = (*s)[:n-1]
}
