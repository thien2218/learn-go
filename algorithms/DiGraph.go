package algorithms

import "log"

type diGraph[V Node] struct {
	graph adjacencyList[V]
}

func NewDiGraph[V Node]() *diGraph[V] {
	mg := new(diGraph[V])
	mg.graph = make(adjacencyList[V])
	return mg
}

func (mg diGraph[V]) GetEdges(vertex V) []Edge[V] {
	checkVertex[V](mg.graph, vertex)
	return mg.graph[vertex]
}

func (mg diGraph[V]) Insert(vertex V, edges ...Edge[V]) {
	insertToGraph[V](mg.graph, vertex, edges...)
}

func (mg diGraph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		mg.graph[vertex][index].Weight = weights[i]
	}
}

func (mg diGraph[V]) Delete(vertex V) {
	deleteVertex[V](mg.graph, vertex)
}
