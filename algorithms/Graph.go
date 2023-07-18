package algorithms

import "log"

type graph[V Node] struct {
	graph adjacencyList[V]
}

func NewGraph[V Node]() *graph[V] {
	g := new(graph[V])
	g.graph = make(adjacencyList[V])
	return g
}

func (g graph[V]) GetEdges(vertex V) []Edge[V] {
	checkVertex[V](g.graph, vertex)
	return g.graph[vertex]
}

func (g graph[V]) Insert(vertex V, edges ...Edge[V]) {
	graph := g.graph
	insertToGraph[V](graph, vertex, edges...)

	for weight, edge := range edges {
		newEdge := new(Edge[V])
		newEdge.EndVertex = vertex
		newEdge.Weight = float64(weight)
		graph[edge.EndVertex] = append(graph[edge.EndVertex], *newEdge)
	}
}

func (g graph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
	if len(edgeIndexes) != len(weights) {
		log.Fatal("Index list size must be the same as weight list!")
	}

	for i, index := range edgeIndexes {
		g.graph[vertex][index].Weight = weights[i]
		otherVertex := g.graph[vertex][index].EndVertex

		for j, edge := range g.graph[otherVertex] {
			if edge.EndVertex == vertex {
				g.graph[otherVertex][j].Weight = weights[i]
			}
		}
	}
}

func (g graph[V]) Delete(vertex V) {
	graph := g.graph

	for _, edge := range graph[vertex] {
		l := len(graph[edge.EndVertex])
		for i := 0; i < l; i++ {
			if graph[edge.EndVertex][i].EndVertex == vertex {
				graph[edge.EndVertex][i], graph[edge.EndVertex][l-1] = graph[edge.EndVertex][l-1], graph[edge.EndVertex][i]
				graph[edge.EndVertex] = graph[edge.EndVertex][:l-1]
				i--
				l--
			}
		}
	}

	deleteVertex[V](graph, vertex)
}
