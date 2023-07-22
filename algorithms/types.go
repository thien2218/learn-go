package algorithms

import "golang.org/x/exp/constraints"

type Node interface {
	constraints.Ordered
}

type Edge[V Node] struct {
	EndVertex V
	Weight    float64
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
