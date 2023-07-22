package algorithms

import "log"

// GRAPH HELPER FUNCTIONS

func checkVertex[V Node](vertices adjacencyList[V], vertex V) {
	if _, exist := vertices[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
}

func deleteVertex[V Node](vertices adjacencyList[V], vertex V) {
	checkVertex[V](vertices, vertex)
	delete(vertices, vertex)
}

func insertToGraph[V Node](vertices map[V][]Edge[V], vertex V, edges ...Edge[V]) {
	if _, exist := vertices[vertex]; !exist {
		vertices[vertex] = edges
	} else {
		vertices[vertex] = append(vertices[vertex], edges...)
	}

	for _, edge := range edges {
		if _, exist := vertices[edge.EndVertex]; !exist {
			vertices[edge.EndVertex] = make([]Edge[V], 0)
		}
	}
}

// CONSTRUCTOR

func NewGraph[V Node](graphType string) *Graph[V] {
	if graphType != "directed" && graphType != "undirected" {
		log.Fatal("Invalid graph type")
	}

	g := new(Graph[V])
	g.vertices = make(adjacencyList[V])
	g.graphType = graphType

	return g
}

// PUBLIC

func (g Graph[V]) GetEdges(vertex V) []Edge[V] {
	checkVertex[V](g.vertices, vertex)
	return g.vertices[vertex]
}

func (g Graph[V]) Insert(vertex V, edges ...Edge[V]) {
	vertices := g.vertices
	insertToGraph[V](vertices, vertex, edges...)

	if g.graphType == "undirected" {
		for _, edge := range edges {
			newEdge := new(Edge[V])
			newEdge.EndVertex = vertex
			newEdge.Weight = float64(edge.Weight)
			vertices[edge.EndVertex] = append(vertices[edge.EndVertex], *newEdge)
		}
	}
}

// func (g Graph[V]) Update(vertex V, edgeIndexes []int, weights []float64) {
// 	if len(edgeIndexes) != len(weights) {
// 		log.Fatal("Index list size must be the same as weight list!")
// 	}

// 	for i, index := range edgeIndexes {
// 		g.graph[vertex][index].Weight = weights[i]
// 		otherVertex := g.graph[vertex][index].EndVertex

// 		for j, edge := range g.graph[otherVertex] {
// 			if edge.EndVertex == vertex {
// 				g.graph[otherVertex][j].Weight = weights[i]
// 			}
// 		}
// 	}
// }

func (g Graph[V]) Delete(vertex V) {
	vertices := g.vertices

	if g.graphType == "undirected" {
		for _, edge := range vertices[vertex] {
			l := len(vertices[edge.EndVertex])
			for i := 0; i < l; i++ {
				if vertices[edge.EndVertex][i].EndVertex == vertex {
					vertices[edge.EndVertex][i], vertices[edge.EndVertex][l-1] = vertices[edge.EndVertex][l-1], vertices[edge.EndVertex][i]
					vertices[edge.EndVertex] = vertices[edge.EndVertex][:l-1]
					i--
					l--
				}
			}
		}
	}

	deleteVertex[V](vertices, vertex)
}
