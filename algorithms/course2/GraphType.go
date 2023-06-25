package algorithms

import "golang.org/x/exp/constraints"

type Vertex interface {
	constraints.Ordered | *interface{}
}

type WeightedEdge[V Vertex] struct {
	endVertex V
	weight    float64
}

type IGraph[V Vertex] interface {
	Insert(vertex V, edges map[V]float64)
	Update(vertex V, edges map[V]float64)
	Delete(vertex V)
}

type IMultiGraph[V Vertex] interface {
	Insert(vertex V, edges ...WeightedEdge[V])
	Update(vertex V, edgeIndexes []int, weights []float64)
	Delete(vertex V)
}

type simple[V Vertex] map[V]map[V]float64
type multi[V Vertex] map[V][]WeightedEdge[V]
