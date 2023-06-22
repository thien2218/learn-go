package algorithms

import "log"

type MultiDiGraph[N Node] struct {
	graph multiGraph[N]
}

func (mdg MultiDiGraph[N]) Insert(vertex N, edges ...WeightedEdge[N]) {
	insertToMGraph[N](mdg.graph, vertex, edges...)
}

func (mdg MultiDiGraph[N]) Update(vertex N, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		mdg.graph[vertex][index].weight = weights[i]
	}
}

func (mdg MultiDiGraph[N]) Delete(vertex N) {
	deleteVertex[N](mdg.graph, vertex)
}
