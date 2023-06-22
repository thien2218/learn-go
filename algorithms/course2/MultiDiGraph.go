package algorithms

import "log"

type MultiDiGraph[N Node] map[N][]WeightedEdge[N]

func (mdg MultiDiGraph[N]) Insert(vertex N, edges ...WeightedEdge[N]) {
	if _, exist := mdg[vertex]; !exist {
		mdg[vertex] = edges
	} else {
		mdg[vertex] = append(mdg[vertex], edges...)
	}

	for _, edge := range edges {
		if _, exist := mdg[edge.endVertex]; !exist {
			mdg[edge.endVertex] = make([]WeightedEdge[N], 0)
		}
	}
}

func (mdg MultiDiGraph[N]) Delete(vertex N) {
	if _, exist := mdg[vertex]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}

	delete(mdg, vertex)
}

func (mdg MultiDiGraph[N]) BFS(start N, target N) bool {
	if _, exist := mdg[start]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}
	if start == target {
		return true
	}

	queue := []N{start}
	visited := map[N]interface{}{start: nil}

	for len(queue) > 0 {
		curr := queue[0]

		for _, edge := range mdg[curr] {
			if edge.endVertex == target {
				return true
			}
			if _, isVisited := visited[edge.endVertex]; !isVisited {
				visited[edge.endVertex] = nil
				queue = append(queue, edge.endVertex)
			}
		}

		queue = queue[1:]
	}

	return false
}

func (mdg MultiDiGraph[N]) dfsLoop(start N, target N, visited map[N]interface{}) bool {
	if start == target {
		return true
	}

	for _, edge := range mdg[start] {
		if _, isVisited := visited[edge.endVertex]; !isVisited {
			visited[edge.endVertex] = nil
			mdg.dfsLoop(edge.endVertex, target, visited)
		}
	}

	return false
}

func (mdg MultiDiGraph[N]) DFS(start N, target N) bool {
	if _, exist := mdg[start]; !exist {
		log.Fatal("Vertex does not exist in graph!")
	}

	visited := map[N]interface{}{start: nil}
	return mdg.dfsLoop(start, target, visited)
}
