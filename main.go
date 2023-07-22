package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thien2218/learn-go/algorithms"
)

func main() {
	file, err := os.Open("excercises/prim-mst.txt")
	handleError(err)
	defer file.Close()

	graph := algorithms.NewGraph[int]("undirected")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		vertex, err1 := strconv.Atoi(nums[0])
		handleError(err1)
		endVertex, err2 := strconv.Atoi(nums[1])
		handleError(err2)
		weight, err3 := strconv.ParseFloat(nums[2], 64)
		handleError(err3)

		edge := algorithms.Edge[int]{EndVertex: endVertex, Weight: weight}

		graph.Insert(vertex, edge)
	}

	cost := algorithms.PrimMst[int](*graph, 1)
	fmt.Println(cost)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
