package algorithms

import "log"

type Graph[N Node] struct {
	graph graph[N]
}

func checkVertex[N Node, E interface{}](graph map[N]E, vertex N) {
	if _, exist := graph[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
}

func deleteVertex[N Node, E interface{}](graph map[N]E, vertex N) {
	checkVertex[N](graph, vertex)
	delete(graph, vertex)
}

func insertToGraph[N Node](graph map[N]map[N]float64, vertex N, edges map[N]float64) {
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	}

	for endVertex, weight := range edges {
		if _, exist := graph[endVertex]; !exist {
			graph[endVertex] = make(map[N]float64)
		}
		if _, exist := graph[vertex][endVertex]; !exist {
			graph[vertex][endVertex] = weight
		}
	}
}

func updateGraph[N Node](graph map[N]map[N]float64, vertex N, edges map[N]float64) {
	checkVertex[N](graph, vertex)

	for endVertex, weight := range edges {
		if _, exist := graph[vertex][endVertex]; exist {
			graph[vertex][endVertex] = weight
		}
	}
}

func (g Graph[N]) Insert(vertex N, edges map[N]float64) {
	graph := g.graph
	insertToGraph[N](g.graph, vertex, edges)

	for endVertex, weight := range edges {
		graph[endVertex][vertex] = weight
	}
}

func (g Graph[N]) Update(vertex N, edges map[N]float64) {
	graph := g.graph
	updateGraph[N](g.graph, vertex, edges)

	for endVertex, weight := range edges {
		if _, exist := g.graph[endVertex]; exist {
			graph[endVertex][vertex] = weight
		}
	}
}

func (g Graph[N]) Delete(vertex N) {
	graph := g.graph

	for endVertex := range graph[vertex] {
		delete(graph[endVertex], vertex)
	}

	deleteVertex[N](g.graph, vertex)
}
