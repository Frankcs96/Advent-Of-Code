package day12

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

var Count int
var Data map[string][]string
var LowerCave string

func Solution(filename string) int {

	paths := GetInput(filename)

	graph := Graph{}

	for _, path := range paths {

		graph.AddNode(path.StartAt)
		graph.AddNode(path.EndAt)

	}

	for _, path := range paths {

		graph.AddConnection(path.StartAt, path.EndAt)
		graph.AddConnection(path.EndAt, path.StartAt)

	}

	Count = 0
	Data = graph
	goneOnce := true
	for _, cave := range Data["start"] {
		visitedLower := []string{"start"}
		Explore(cave, visitedLower, goneOnce)
	}
	return Count

}

func SolutionPartTwo(filename string) int {
	paths := GetInput(filename)

	graph := Graph{}

	for _, path := range paths {

		graph.AddNode(path.StartAt)
		graph.AddNode(path.EndAt)

	}

	for _, path := range paths {

		graph.AddConnection(path.StartAt, path.EndAt)
		graph.AddConnection(path.EndAt, path.StartAt)

	}

	finalCount := 0
	Data = graph
	lowerList := MakeLowerList(graph)
	mainPathCount := Solution(filename)
	finalCount += mainPathCount
	for _, lowerCave := range lowerList {
		LowerCave = lowerCave
		Count = 0
		goneOnce := false
		for _, cave := range Data["start"] {
			visitedLower := []string{"start"}
			Explore(cave, visitedLower, goneOnce)
		}
		thisCount := Count - mainPathCount
		finalCount += thisCount

	}

	return finalCount
}

func GetInput(filename string) []Path {

	input := []Path{}
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val := scanner.Text()
		splitPath := strings.Split(val, "-")
		path := Path{StartAt: splitPath[0], EndAt: splitPath[1]}
		input = append(input, path)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func MakeLowerList(data map[string][]string) []string {
	lowerList := []string{}
	for key := range data {
		if key != "start" && key != "end" && !IsUpper(key) {
			lowerList = append(lowerList, key)
		}
	}
	return lowerList
}

func IsUpper(chars string) bool {
	return unicode.IsUpper(rune(chars[0]))
}

func NotVisited(cave string, visitedLower []string) bool {
	for _, lower := range visitedLower {
		if cave == lower {
			return false
		}
	}
	return true
}

func Explore(thisCave string, visitedLower []string, goneOnce bool) {
	if thisCave == "end" {
		Count++
		return
	} else {

		if !IsUpper(thisCave) {
			if goneOnce {
				visitedLower = append(visitedLower, thisCave)
			} else {
				if thisCave == LowerCave {
					goneOnce = true
				} else {
					visitedLower = append(visitedLower, thisCave)
				}
			}
		}

		for _, cave := range Data[thisCave] {
			if NotVisited(cave, visitedLower) {
				Explore(cave, visitedLower, goneOnce)
			}

		}
	}
}

type Path struct {
	StartAt string
	EndAt   string
}

type Graph map[string][]string

func (graph Graph) AddNode(node string) {

	graph[node] = nil
}

func (graph Graph) AddConnection(value string, way string) {

	graph[value] = append(graph[value], way)
}
