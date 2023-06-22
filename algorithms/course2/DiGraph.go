package algorithms

import "log"

type DiGraph[N Node] map[N]map[N]float64

func (dg DiGraph[N]) Insert(vertex N, edges map[N]float64) {
	insertToGraph[N](dg, vertex, edges)
}

func (dg DiGraph[N]) Delete(vertex N) {
	deleteVertex[N](dg, vertex)
}

func (dg DiGraph[N]) BFS(start N, target N) bool {
	if _, exist := dg[start]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
	if start == target {
		return true
	}

	queue := []N{start}
	visited := map[N]interface{}{start: nil}

	for len(queue) > 0 {
		curr := queue[0]

		for endVertex := range dg[curr] {
			if endVertex == target {
				return true
			}
			if _, isVisited := visited[endVertex]; !isVisited {
				visited[endVertex] = nil
				queue = append(queue, endVertex)
			}
		}

		queue = queue[1:]
	}

	return false
}

func (dg DiGraph[N]) dfsLoop(start N, target N, visited map[N]interface{}) bool {
	if start == target {
		return true
	}

	for endVertex := range dg[start] {
		if _, isVisited := visited[endVertex]; !isVisited {
			visited[endVertex] = 0
			dg.dfsLoop(endVertex, target, visited)
		}
	}

	return false
}

func (dg DiGraph[N]) DFS(start N, target N) bool {
	if _, exist := dg[start]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}

	visited := map[N]interface{}{start: nil}
	return dg.dfsLoop(start, target, visited)
}
