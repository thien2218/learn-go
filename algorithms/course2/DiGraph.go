package algorithms

import "log"

type DiGraph[V Vertex] struct {
	graph adjacencyList[V]
}

func NewMultiDiGraph[V Vertex]() *DiGraph[V] {
	mg := new(DiGraph[V])
	mg.graph = make(adjacencyList[V])
	return mg
}

func (mg DiGraph[V]) Insert(vertex V, edges ...Edge[V]) {
	insertToGraph[V](mg.graph, vertex, edges...)
}

func (mg DiGraph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		mg.graph[vertex][index].weight = weights[i]
	}
}

func (mg DiGraph[V]) Delete(vertex V) {
	deleteVertex[V](mg.graph, vertex)
}
