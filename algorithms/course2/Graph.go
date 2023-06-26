package algorithms

import "log"

type Graph[V Vertex] struct {
	graph adjacencyList[V]
}

func NewMultiGraph[V Vertex]() *Graph[V] {
	g := new(Graph[V])
	g.graph = make(adjacencyList[V])
	return g
}

func (g Graph[V]) Insert(vertex V, edges ...Edge[V]) {
	graph := g.graph
	insertToGraph[V](graph, vertex, edges...)

	for weight, edge := range edges {
		newEdge := new(Edge[V])
		newEdge.endVertex = vertex
		newEdge.weight = float64(weight)
		graph[edge.endVertex] = append(graph[edge.endVertex], *newEdge)
	}
}

func (g Graph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		g.graph[vertex][index].weight = weights[i]
		otherVertex := g.graph[vertex][index].endVertex

		for j, edge := range g.graph[otherVertex] {
			if edge.endVertex == vertex {
				g.graph[otherVertex][j].weight = weights[i]
			}
		}
	}
}

func (g Graph[V]) Delete(vertex V) {
	graph := g.graph

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
