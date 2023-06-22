package algorithms

import (
	"log"

	"golang.org/x/exp/constraints"
)

type Node interface {
	constraints.Ordered | *interface{}
}

type WeightedEdge[N Node] struct {
	endVertex N
	weight    float64
}

func deleteVertex[N Node, E interface{}](graph map[N]E, vertex N) {
	if _, exist := graph[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}

	delete(graph, vertex)
}

func insertToGraph[N Node](graph map[N]map[N]float64, vertex N, edges map[N]float64) {
	if _, exist := graph[vertex]; !exist {
		graph[vertex] = edges
	}

	for endVertex := range edges {
		if _, exist := graph[endVertex]; !exist {
			graph[endVertex] = make(map[N]float64)
		}
		if _, exist := graph[vertex][endVertex]; !exist {
			graph[vertex][endVertex] = 0
		}
	}
}

type Graph[N Node] DiGraph[N]

func (g Graph[N]) Insert(vertex N, edges map[N]float64) {
	insertToGraph[N](g, vertex, edges)

	for endVertex := range edges {
		g[endVertex][vertex] = 0
	}
}

func (g Graph[N]) Delete(vertex N) {
	graph := g

	for endVertex := range graph[vertex] {
		delete(graph[endVertex], vertex)
	}

	deleteVertex[N](g, vertex)
}
