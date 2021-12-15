package day15

import (
	"bufio"
	"container/heap"
	"log"
	"os"
)

func Solution(filename string) int {
	grid := GetInput(filename)

	risks := Dijkstra(grid)

	return risks
}

func SolutionPartTwo(filename string) int {

	tile := GetInput(filename)

	fullMap := CalculateFullMap(tile)
	risks := Dijkstra(fullMap)
	return risks

}

// Dijkstra First time implementing this algorithm so dunno if it's the best way, but was fun!
func Dijkstra(grid [][]int) int {
	nodesPq := &NodePq{}
	heap.Init(nodesPq)

	visited := make(map[Point]bool)

	heap.Push(nodesPq, Node{Risk: 0, Position: Point{I: 0, J: 0}})

	for nodesPq.Len() > 0 {

		currentNode := heap.Pop(nodesPq).(Node)

		if _, ok := visited[currentNode.Position]; ok {

			continue

		}
		if currentNode.Position.I == len(grid)-1 && currentNode.Position.J == len(grid)-1 {

			return currentNode.Risk
		}

		i := currentNode.Position.I
		j := currentNode.Position.J

		visited[Point{I: i, J: j}] = true
		if i-1 >= 0 {
			topNodeRisk := grid[i-1][j]

			topNode := Node{Risk: currentNode.Risk + topNodeRisk, Position: Point{I: i - 1, J: j}}

			heap.Push(nodesPq, topNode)
		}
		if j+1 < len(grid) {
			rightNodeRisk := grid[i][j+1]
			rightNode := Node{Risk: currentNode.Risk + rightNodeRisk, Position: Point{I: i, J: j + 1}}

			heap.Push(nodesPq, rightNode)

		}

		if i+1 < len(grid) {
			botNodeRisk := grid[i+1][j]
			botNode := Node{Risk: currentNode.Risk + botNodeRisk, Position: Point{I: i + 1, J: j}}

			heap.Push(nodesPq, botNode)
		}

		if j-1 >= 0 {
			leftNodeRisk := grid[i][j-1]
			leftNode := Node{Risk: currentNode.Risk + leftNodeRisk, Position: Point{I: i, J: j - 1}}

			heap.Push(nodesPq, leftNode)
		}

	}
	return -1
}

func CalculateFullMap(tile [][]int) [][]int {
	var grid = make([][]int, len(tile)*5)

	for i := range grid {
		grid[i] = make([]int, len(tile)*5)
	}

	for i := 0; i < len(tile); i++ {
		for j := 0; j < len(tile); j++ {

			grid[i][j] = tile[i][j]
			for k := 1; k < 5; k++ {
				if tile[i][j]+k > 9 {
					grid[i][j+len(tile)*k] = tile[i][j] + k - 9
				} else {
					grid[i][j+len(tile)*k] = tile[i][j] + k
				}
			}
		}
	}
	for i := len(tile); i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i-len(tile)][j]+1 > 9 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i-len(tile)][j] + 1
			}
		}
	}
	return grid
}

func GetInput(filename string) [][]int {

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

	var intInput = make([][]int, len(input))
	for i := range input {
		intInput[i] = make([]int, len(input))
	}

	for i := 0; i < len(input); i++ {

		for j := 0; j < len(input[0]); j++ {

			intInput[i][j] = int(input[i][j] - 48)

		}
	}

	return intInput
}

type Point struct {
	I, J int
}

type Node struct {
	Position Point
	Risk     int
}
type NodePq []Node

//Methods needed for the interface in the heap

func (priorityQueue NodePq) Len() int { return len(priorityQueue) }
func (priorityQueue NodePq) Less(i, j int) bool {
	return priorityQueue[i].Risk < priorityQueue[j].Risk
}
func (priorityQueue NodePq) Swap(i, j int) {
	priorityQueue[i], priorityQueue[j] = priorityQueue[j], priorityQueue[i]
}
func (priorityQueue *NodePq) Push(x interface{}) { *priorityQueue = append(*priorityQueue, x.(Node)) }
func (priorityQueue *NodePq) Pop() (popped interface{}) {
	popped = (*priorityQueue)[len(*priorityQueue)-1]
	*priorityQueue = (*priorityQueue)[:len(*priorityQueue)-1]
	return
}
