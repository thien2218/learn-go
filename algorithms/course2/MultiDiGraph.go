package algorithms

import "log"

type multiDiGraph[V Vertex] struct {
	graph multi[V]
}

func NewMultiDiGraph[V Vertex]() *multiDiGraph[V] {
	mdg := new(multiDiGraph[V])
	mdg.graph = make(multi[V])
	return mdg
}

func (mdg multiDiGraph[V]) Insert(vertex V, edges ...WeightedEdge[V]) {
	insertToMGraph[V](mdg.graph, vertex, edges...)
}

func (mdg multiDiGraph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		mdg.graph[vertex][index].weight = weights[i]
	}
}

func (mdg multiDiGraph[V]) Delete(vertex V) {
	deleteVertex[V](mdg.graph, vertex)
}
