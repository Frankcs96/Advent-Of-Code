package day08

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution(filename string) int {
	entries := GetInput(filename)

	solution := 0

	for _, entry := range entries {

		outputs := strings.Split(entry.Output, " ")

		for _, output := range outputs {
			if len(output) == 2 || len(output) == 4 || len(output) == 7 || len(output) == 3 {

				solution++

			}

		}
	}
	return solution
}

func SolutionPartTwo(filename string) int {

	entries := GetInput(filename)
	signalPatterns := []string{}
	outputs := []string{}
	// m := map[string]int{}

	sum := 0
	for _, entry := range entries {

		signalPatterns = strings.Split(entry.SignalPattern, " ")

		outputs = strings.Split(entry.Output, " ")
		topWire := CalculateTopWire(signalPatterns)
		bottomWire := CalculateBottomWire(signalPatterns, topWire)
		bottomLeftWire := CalculateBotomLeftWire(signalPatterns, bottomWire, topWire)
		upperRightWire := calculateTopRight(signalPatterns)
		bottomRightWire := calculateBottomRightWire(signalPatterns, upperRightWire)
		middleWire := calculateMiddleWire(signalPatterns, upperRightWire, bottomLeftWire)
		upperLeftWire := calculateUpperLeftWire(topWire, bottomWire, bottomLeftWire, upperRightWire, bottomRightWire, middleWire)

		wiredNumbers := WiredNumber{

			Zero:  topWire + upperRightWire + bottomRightWire + bottomWire + bottomLeftWire + upperLeftWire,
			One:   upperRightWire + bottomRightWire,
			Two:   topWire + upperRightWire + middleWire + bottomLeftWire + bottomWire,
			Three: topWire + upperRightWire + middleWire + bottomRightWire + bottomWire,
			Four:  upperLeftWire + upperRightWire + middleWire + bottomRightWire,
			Five:  topWire + upperLeftWire + middleWire + bottomRightWire + bottomWire,
			Six:   topWire + upperLeftWire + middleWire + bottomLeftWire + bottomRightWire + bottomWire,
			Seven: topWire + upperRightWire + bottomRightWire,
			Eight: "abcdefg",
			Nine:  topWire + upperLeftWire + upperRightWire + middleWire + bottomRightWire + bottomWire,
		}

		solution := ""

		for _, v := range outputs {
			sortedOutput := SortString(v)

			if sortedOutput == SortString(wiredNumbers.Zero) {

				solution += "0"
			}
			if sortedOutput == SortString(wiredNumbers.One) {

				solution += "1"
			}
			if sortedOutput == SortString(wiredNumbers.Two) {

				solution += "2"
			}
			if sortedOutput == SortString(wiredNumbers.Three) {

				solution += "3"
			}
			if sortedOutput == SortString(wiredNumbers.Four) {

				solution += "4"
			}
			if sortedOutput == SortString(wiredNumbers.Five) {

				solution += "5"
			}
			if sortedOutput == SortString(wiredNumbers.Six) {

				solution += "6"
			}
			if sortedOutput == SortString(wiredNumbers.Seven) {

				solution += "7"
			}
			if sortedOutput == SortString(wiredNumbers.Eight) {

				solution += "8"
			}
			if sortedOutput == SortString(wiredNumbers.Nine) {

				solution += "9"
			}

		}

		number, err := strconv.Atoi(solution)

		if err != nil {
			log.Fatal("cannot cast string to int")
		}

		sum += number

	}

	return sum
}

func calculateUpperLeftWire(topWire, bottomWire, bottomLeftWire, upperRightWire, bottomRightWire, middleWire string) string {
	var letters = []string{"a", "b", "c", "d", "e", "f", "g"}

	solution := ""
	for _, v := range letters {

		if v != topWire && v != bottomWire && v != bottomLeftWire && v != upperRightWire && v != bottomRightWire && v != middleWire {

			solution = v

		}

	}

	return solution

}

func calculateMiddleWire(signalPatterns []string, upperRightWire, bottomLeftWire string) string {

	var numbersWithSixWires []string
	var wiresOfEight string
	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 6 {

			numbersWithSixWires = append(numbersWithSixWires, signalPattern)

		}
		if len(signalPattern) == 7 {

			wiresOfEight = signalPattern

		}
	}
	numberZeroWires := ""
	for _, numberWithSixWire := range numbersWithSixWires {

		if strings.ContainsAny(numberWithSixWire, upperRightWire) && strings.ContainsAny(numberWithSixWire, bottomLeftWire) {

			numberZeroWires = numberWithSixWire
		}

	}

	solution := ""
	for _, runeOfEight := range wiresOfEight {

		if !strings.ContainsRune(numberZeroWires, runeOfEight) {

			solution = string(runeOfEight)

		}
	}

	return solution
}

func calculateBottomRightWire(signalPatterns []string, upperRightWire string) string {
	var wiresForOne []string

	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 2 {

			wiresForOne = strings.Split(signalPattern, "")

		}

	}

	solution := ""
	for _, v := range wiresForOne {

		if v != upperRightWire {

			solution = v
		}

	}
	return solution
}

func calculateTopRight(signalPatterns []string) string {

	var numbersWithSixWires []string

	var wiresForOne []string

	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 2 {

			wiresForOne = strings.Split(signalPattern, "")

		}

		if len(signalPattern) == 6 {

			numbersWithSixWires = append(numbersWithSixWires, signalPattern)

		}
	}

	solution := ""
	for _, numberWithSixWire := range numbersWithSixWires {

		for _, wireOfOne := range wiresForOne {

			if !strings.ContainsAny(numberWithSixWire, wireOfOne) {

				solution = wireOfOne
			}
		}

	}
	return solution
}

func CalculateBotomLeftWire(signalPatterns []string, bottomWire string, topWire string) string {
	var wiresForEight string
	var wiresForNine string

	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 7 {

			wiresForEight = signalPattern

		}
		if len(signalPattern) == 4 {

			wiresForNine = signalPattern + bottomWire + topWire

		}

	}
	s1 := 0
	s2 := 0
	for _, v := range wiresForEight {
		s1 += int(v)
	}
	for _, v := range wiresForNine {
		s2 += int(v)
	}
	difference := s1 - s2
	rune := rune(difference)
	return string(rune)
}

func CalculateBottomWire(signalPatterns []string, topWire string) string {
	var wiresForFour string
	var numbersWithSixWires []string
	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 4 {

			wiresForFour = signalPattern

		}
		if len(signalPattern) == 6 {

			numbersWithSixWires = append(numbersWithSixWires, signalPattern)

		}
	}
	wiresOfFourAndTop := wiresForFour + topWire

	wiresForNine := ""
	for _, numberWithSixWires := range numbersWithSixWires {
		containsAllRunes := true
		for _, r := range wiresOfFourAndTop {

			if !strings.ContainsRune(numberWithSixWires, r) {

				containsAllRunes = false
			}

		}
		if containsAllRunes {

			wiresForNine = numberWithSixWires
		}
	}

	s1 := 0
	s2 := 0
	for _, v := range wiresForNine {
		s1 += int(v)
	}
	for _, v := range wiresOfFourAndTop {
		s2 += int(v)
	}
	difference := s1 - s2
	rune := rune(difference)
	return string(rune)
}

func CalculateTopWire(signalPatterns []string) string {

	var wiresForOne string
	var wiresForSeven string
	for _, signalPattern := range signalPatterns {

		if len(signalPattern) == 2 {

			wiresForOne = signalPattern
		}
		if len(signalPattern) == 3 {

			wiresForSeven = signalPattern
		}
	}

	s1 := 0
	s2 := 0
	for _, v := range wiresForOne {
		s1 += int(v)
	}
	for _, v := range wiresForSeven {
		s2 += int(v)
	}
	difference := s2 - s1
	rune := rune(difference)
	return string(rune)
}

func GetInput(filename string) []Entry {

	input := []string{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()
		parsedString := strings.Split(val, " | ")
		input = append(input, parsedString...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	entries := []Entry{}

	for i := 0; i < len(input)-1; i += 2 {

		entry := Entry{

			SignalPattern: input[i],
			Output:        input[i+1],
		}

		entries = append(entries, entry)

	}
	return entries
}

type Entry struct {
	SignalPattern string
	Output        string
}

type WiredNumber struct {
	Zero  string
	One   string
	Two   string
	Three string
	Four  string
	Five  string
	Six   string
	Seven string
	Eight string
	Nine  string
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
