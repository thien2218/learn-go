package algorithms

import "golang.org/x/exp/constraints"

type Node interface {
	constraints.Ordered
}

type Edge[V Node] struct {
	EndVertex V
	Weight    float64
}

type IGraph[V Node] interface {
	GetEdges(vertex V) []Edge[V]
	Insert(vertex V, edges ...Edge[V])
	Update(vertex V, edgeIndexes []int, weights []float64)
	Delete(vertex V)
}

type adjacencyList[V Node] map[V][]Edge[V]

type Graph[V Node] struct {
	vertices  adjacencyList[V]
	graphType string
}

type Heap[N Node] struct {
	nodes    []N
	heapType string
}
