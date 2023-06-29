package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	algorithms "github.com/thien2218/learn-go/algorithms/course2"
)

func main() {
	file, err := os.Open("excercises/median.txt")
	handleError(err)
	defer file.Close()

	arr := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		handleError(err)
		arr = append(arr, num)
	}

	algorithms.Median(arr)
}

// curr := "1"
// edges := make(map[int]float64)
// diGraph := algorithms.NewDiGraph[int]()

// file, err := os.Open("scc.txt")
// handleError(err)
// defer file.Close()

// scanner := bufio.NewScanner(file)

// for scanner.Scan() {
// 	line := strings.Split(strings.TrimSpace(scanner.Text()), " ")

// 	if line[0] != curr {
// 		vertex, err := strconv.Atoi(curr)
// 		handleError(err)

// 		diGraph.Insert(vertex, edges)

// 		curr = line[0]
// 		edges = make(map[int]float64)
// 	}

// 	EndVertex, err := strconv.Atoi(line[1])
// 	handleError(err)

// 	edges[EndVertex] = 0
// }

// vertex, err := strconv.Atoi(curr)
// handleError(err)

// diGraph.Insert(vertex, edges)

// algorithms.ComputeScc[int](diGraph)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
