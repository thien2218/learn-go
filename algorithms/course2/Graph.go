package algorithms

type graph[V Vertex] struct {
	graph simple[V]
}

func NewGraph[V Vertex]() *graph[V] {
	g := new(graph[V])
	g.graph = make(simple[V])
	return g
}

func (g graph[V]) Insert(vertex V, edges map[V]float64) {
	graph := g.graph
	insertToGraph[V](g.graph, vertex, edges)

	for endVertex, weight := range edges {
		graph[endVertex][vertex] = weight
	}
}

func (g graph[V]) Update(vertex V, edges map[V]float64) {
	graph := g.graph
	updateGraph[V](g.graph, vertex, edges)

	for endVertex, weight := range edges {
		if _, exist := g.graph[endVertex]; exist {
			graph[endVertex][vertex] = weight
		}
	}
}

func (g graph[V]) Delete(vertex V) {
	graph := g.graph

	for endVertex := range graph[vertex] {
		delete(graph[endVertex], vertex)
	}

	deleteVertex[V](g.graph, vertex)
}
