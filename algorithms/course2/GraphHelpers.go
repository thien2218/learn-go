package algorithms

import "log"

func checkVertex[V Vertex, E interface{}](graph map[V]E, vertex V) {
	if _, exist := graph[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
}

func deleteVertex[V Vertex, E interface{}](graph map[V]E, vertex V) {
	checkVertex[V](graph, vertex)
	delete(graph, vertex)
}

func insertToGraph[V Vertex](graph map[V]map[V]float64, vertex V, edges map[V]float64) {
	// Initialize vertex if it doesn't exist
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	}

	for endVertex, weight := range edges {
		// Initialize other vertices at the other end of an edge
		// if it doesn't exist
		if _, exist := graph[endVertex]; !exist {
			graph[endVertex] = make(map[V]float64)
		}
		// Append an edge if it doesn't exist
		if _, exist := graph[vertex][endVertex]; !exist {
			graph[vertex][endVertex] = weight
		}
	}
}

func updateGraph[V Vertex](graph map[V]map[V]float64, vertex V, edges map[V]float64) {
	checkVertex[V](graph, vertex)

	for endVertex, weight := range edges {
		if _, exist := graph[vertex][endVertex]; exist {
			graph[vertex][endVertex] = weight
		}
	}
}

func insertToMGraph[V Vertex](graph map[V][]WeightedEdge[V], vertex V, edges ...WeightedEdge[V]) {
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	} else {
		graph[vertex] = append(graph[vertex], edges...)
	}

	for _, edge := range edges {
		if _, exist := graph[edge.endVertex]; !exist {
			graph[edge.endVertex] = make([]WeightedEdge[V], 0)
		}
	}
}
