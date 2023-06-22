package algorithms

type DiGraph[N Node] struct {
	graph graph[N]
}

func (dg DiGraph[N]) Insert(vertex N, edges map[N]float64) {
	insertToGraph[N](dg.graph, vertex, edges)
}

func (dg DiGraph[N]) Update(vertex N, edges map[N]float64) {
	updateGraph[N](dg.graph, vertex, edges)
}

func (dg DiGraph[N]) Delete(vertex N) {
	deleteVertex[N](dg.graph, vertex)
}
