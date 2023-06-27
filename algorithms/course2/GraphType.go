package algorithms

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Vertex interface {
	constraints.Ordered | *interface{}
}

type Edge[V Vertex] struct {
	EndVertex V
	Weight    float64
}

type IGraph[V Vertex] interface {
	GetEdges(vertex V) []Edge[V]
	Insert(vertex V, edges ...Edge[V])
	Update(vertex V, edgeIndexes []int, weights []float64)
	Delete(vertex V)
}

type adjacencyList[V Vertex] map[V][]Edge[V]

func checkVertex[V Vertex](graph adjacencyList[V], vertex V) {
	if _, exist := graph[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
}

func deleteVertex[V Vertex](graph adjacencyList[V], vertex V) {
	checkVertex[V](graph, vertex)
	delete(graph, vertex)
}

func insertToGraph[V Vertex](graph map[V][]Edge[V], vertex V, edges ...Edge[V]) {
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	} else {
		graph[vertex] = append(graph[vertex], edges...)
	}

	for _, edge := range edges {
		if _, exist := graph[edge.EndVertex]; !exist {
			graph[edge.EndVertex] = make([]Edge[V], 0)
		}
	}
}
