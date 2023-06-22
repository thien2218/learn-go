package algorithms

import "log"

type WeightedEdge[N Node] struct {
	endVertex N
	weight    float64
}

type MultiGraph[N Node] MultiDiGraph[N]

func insertToMGraph[N Node](graph map[N][]WeightedEdge[N], vertex N, edges ...WeightedEdge[N]) {
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	} else {
		graph[vertex] = append(graph[vertex], edges...)
	}

	for _, edge := range edges {
		if _, exist := graph[edge.endVertex]; !exist {
			graph[edge.endVertex] = make([]WeightedEdge[N], 0)
		}
	}
}

func (mg MultiGraph[N]) Insert(vertex N, edges ...WeightedEdge[N]) {
	graph := mg.graph
	insertToMGraph[N](graph, vertex, edges...)

	for weight, edge := range edges {
		newEdge := new(WeightedEdge[N])
		newEdge.endVertex = vertex
		newEdge.weight = float64(weight)
		graph[edge.endVertex] = append(graph[edge.endVertex], *newEdge)
	}
}

func (mg MultiGraph[N]) Update(vertex N, edgeIndexes []int, weights []float64) {
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

func (mg MultiGraph[N]) Delete(vertex N) {
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

	deleteVertex[N](graph, vertex)
}
