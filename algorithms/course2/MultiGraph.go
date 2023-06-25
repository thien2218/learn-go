package algorithms

import "log"

type multiGraph[V Vertex] struct {
	graph multi[V]
}

func NewMultiGraph[V Vertex]() *multiGraph[V] {
	mg := new(multiGraph[V])
	mg.graph = make(multi[V])
	return mg
}

func (mg multiGraph[V]) Insert(vertex V, edges ...WeightedEdge[V]) {
	graph := mg.graph
	insertToMGraph[V](graph, vertex, edges...)

	for weight, edge := range edges {
		newEdge := new(WeightedEdge[V])
		newEdge.endVertex = vertex
		newEdge.weight = float64(weight)
		graph[edge.endVertex] = append(graph[edge.endVertex], *newEdge)
	}
}

func (mg multiGraph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		mg.graph[vertex][index].weight = weights[i]
		otherVertex := mg.graph[vertex][index].endVertex

		for j, edge := range mg.graph[otherVertex] {
			if edge.endVertex == vertex {
				mg.graph[otherVertex][j].weight = weights[i]
			}
		}
	}
}

func (mg multiGraph[V]) Delete(vertex V) {
	graph := mg.graph

	for _, edge := range graph[vertex] {
		l := len(graph[edge.endVertex])
		for i := 0; i < l; i++ {
			if graph[edge.endVertex][i].endVertex == vertex {
				graph[edge.endVertex][i], graph[edge.endVertex][l-1] = graph[edge.endVertex][l-1], graph[edge.endVertex][i]
				graph[edge.endVertex] = graph[edge.endVertex][:l-1]
				i--
				l--
			}
		}
	}

	deleteVertex[V](graph, vertex)
}
