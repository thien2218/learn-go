package algorithms

import "golang.org/x/exp/constraints"

// Assume that most graphs use in the course are parse graph,
// we'll use Adjacent List to present a graph

// Ex:
// "A": ["B", "C"]
// "B": ["A", "D"]
// "C": ["A", "D"]
// "D": ["B", "C"]

type IGraph interface {
	InsertEdges(node string, connections ...string)
	InsertNode(node string)
	BFS(startNode string, target string)
	DFS(startNode string, target string)
}

type Graph[T constraints.Ordered | *interface{}] map[T][]T
