package algorithms

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
	insertToMGraph[N](mg, vertex, edges...)

	for weight, edge := range edges {
		newEdge := new(WeightedEdge[N])
		newEdge.endVertex = vertex
		newEdge.weight = float64(weight)
		mg[edge.endVertex] = append(mg[edge.endVertex], *newEdge)
	}
}

func (mg MultiGraph[N]) Delete(vertex N) {
	for _, edge := range mg[vertex] {
		l := len(mg[edge.endVertex])
		for i := 0; i < l; i++ {
			if mg[edge.endVertex][i].endVertex == vertex {
				mg[edge.endVertex][i], mg[edge.endVertex][l-1] = mg[edge.endVertex][l-1], mg[edge.endVertex][i]
				mg[edge.endVertex] = mg[edge.endVertex][:l-1]
				i--
				l--
			}
		}
	}

	deleteVertex[N](mg, vertex)
}
