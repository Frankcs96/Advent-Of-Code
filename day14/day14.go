package day14

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func Solution(filename string) int {

	template, rules := GetInput(filename)

	for i := 0; i < 10; i++ {
		template = step(template, rules)
	}

	var counter = make(map[string]int)

	for _, char := range template {
		counter[string(char)] += 1
	}
	var occurrences []int

	for _, times := range counter {
		occurrences = append(occurrences, times)
	}
	sort.Ints(occurrences)

	return occurrences[len(occurrences)-1] - occurrences[0]
}

func SolutionPartTwo(filename string) int64 {
	template, rules := GetInput(filename)
	occurrencesMap := stepOptimized(template, rules, 40)

	occurrences := make([]int64, 0)

	for _, value := range occurrencesMap {
		occurrences = append(occurrences, value)
	}

	sort.Slice(occurrences, func(i, j int) bool { return occurrences[i] < occurrences[j] })
	return occurrences[len(occurrences)-1] - occurrences[0]
}

func stepOptimized(template string, rules map[string]string, times int) map[string]int64 {

	var pairCounter = make(map[string]int64)
	var letterCounter = make(map[string]int64)

	for _, val := range template {
		letterCounter[string(val)] += 1
	}
	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		pairCounter[pair] += 1
	}

	for i := 0; i < times; i++ {

		var m = make(map[string]int64)

		for pair, values := range pairCounter {
			m[pair] = values
		}

		for k := range pairCounter {
			delete(pairCounter, k)
		}

		for pair, val := range m {

			insertion := rules[pair]
			letterCounter[insertion] += val
			firstPair := string(pair[0]) + insertion
			secondPair := insertion + string(pair[1])

			pairCounter[firstPair] += val
			pairCounter[secondPair] += val

		}

	}

	return letterCounter
}

// Slow solution that works in example 1
func step(template string, rules map[string]string) string {
	result := string(template[0])

	for i := 0; i < len(template)-1; i++ {
		finalString := ""
		pair := string(template[i]) + string(template[i+1])

		if val, ok := rules[pair]; ok {
			finalString = val + string(pair[1])

			result += finalString
		}

	}

	return result
}

func GetInput(filename string) (string, map[string]string) {

	var input []string
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

	template := input[0]
	pairs := input[2:]

	var rules = make(map[string]string)

	for _, pair := range pairs {
		splitString := strings.Split(pair, " -> ")
		rules[splitString[0]] = splitString[1]
	}

	return template, rules
}
