package algorithms

type diGraph[V Vertex] struct {
	graph simple[V]
}

func NewDiGraph[V Vertex]() diGraph[V] {
	dg := new(diGraph[V])
	dg.graph = make(simple[V])
	return *dg
}

func (dg diGraph[V]) Insert(vertex V, edges map[V]float64) {
	insertToGraph[V](dg.graph, vertex, edges)
}

func (dg diGraph[V]) Update(vertex V, edges map[V]float64) {
	updateGraph[V](dg.graph, vertex, edges)
}

func (dg diGraph[V]) Delete(vertex V) {
	deleteVertex[V](dg.graph, vertex)
}
