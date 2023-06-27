package main

import (
	"log"
)

func main() {

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
