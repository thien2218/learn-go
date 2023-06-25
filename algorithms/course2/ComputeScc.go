package algorithms

import "fmt"

// The code works but currently is just a piece of crap 'cause
// all building blocks are tangled up and performance is shit.
// TODO: improve the code and performance

func ComputeScc[V Vertex](dg diGraph[V]) {
	dgRev := NewDiGraph[V]()

	visited := make(map[V]interface{})
	time := 0
	ftime := make(map[int]V)
	var count int
	counts := make([]int, 0)

	// 1st pass: Create a reverse version of the original graph
	// whilst computing the finishing times of each vertex
	for vertex, edges := range dg.graph {
		for edge := range edges {
			dgRev.Insert(edge, map[V]float64{vertex: 0})
		}

		if _, isVisited := visited[vertex]; !isVisited {
			dfsRecurse[V](dg, vertex, visited, &time, &count, ftime)
		}
	}

	visited = make(map[V]interface{})

	// 2nd pass: Count the size of each SCC and store it in an array
	for i := time; i > 0; i-- {
		vertex := ftime[i]

		if _, isVisited := visited[vertex]; !isVisited {
			count = 0
			dfsRecurse[V](dgRev, vertex, visited, &time, &count, map[int]V{})
			counts = append(counts, count)
		}
	}

	fmt.Println(counts)
}

func dfsRecurse[V Vertex](dg diGraph[V], vertex V, visited map[V]interface{}, time, count *int, ftime map[int]V) {
	visited[vertex] = nil

	for endVertex := range dg.graph[vertex] {
		if _, isVisited := visited[endVertex]; !isVisited {
			dfsRecurse[V](dg, endVertex, visited, time, count, ftime)
		}
	}

	*count++
	*time++
	ftime[*time] = vertex
}
