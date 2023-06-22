package algorithms

import "golang.org/x/exp/constraints"

type Node interface {
	constraints.Ordered | *interface{}
}

type IGraphCommon[N Node] interface {
	Delete(vertex N)
	BFS(start N, target N) bool
	DFS(start N, target N) bool
}

type IGraph[N Node] interface {
	Insert(vertex N, edges map[N]float64)
	Update(vertex N, edges map[N]float64)
}

type IMultiGraph[N Node] interface {
	Insert(vertex N, edges ...WeightedEdge[N])
	Update(vertex N, edgeIndexes []int, weights []float64)
}

type graph[N Node] map[N]map[N]float64

func (g graph[N]) BFS(start N, target N) bool {
	if start == target {
		return true
	}
	checkVertex[N](g, start)
	checkVertex[N](g, target)

	queue := []N{start}
	visited := map[N]interface{}{start: nil}

	for len(queue) > 0 {
		curr := queue[0]

		for endVertex := range g[curr] {
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

func (g graph[N]) dfsLoop(start N, target N, visited map[N]interface{}) bool {
	if start == target {
		return true
	}

	for endVertex := range g[start] {
		if _, isVisited := visited[endVertex]; !isVisited {
			visited[endVertex] = 0
			g.dfsLoop(endVertex, target, visited)
		}
	}

	return false
}

func (g graph[N]) DFS(start N, target N) bool {
	checkVertex[N](g, start)
	checkVertex[N](g, target)
	visited := map[N]interface{}{start: nil}
	return g.dfsLoop(start, target, visited)
}

type multiGraph[N Node] map[N][]WeightedEdge[N]

func (mg multiGraph[N]) BFS(start N, target N) bool {
	if start == target {
		return true
	}
	checkVertex[N](mg, start)
	checkVertex[N](mg, target)

	queue := []N{start}
	visited := map[N]interface{}{start: nil}

	for len(queue) > 0 {
		curr := queue[0]

		for _, edge := range mg[curr] {
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

func (mg multiGraph[N]) dfsLoop(start N, target N, visited map[N]interface{}) bool {
	if start == target {
		return true
	}

	for _, edge := range mg[start] {
		if _, isVisited := visited[edge.endVertex]; !isVisited {
			visited[edge.endVertex] = nil
			mg.dfsLoop(edge.endVertex, target, visited)
		}
	}

	return false
}

func (mg multiGraph[N]) DFS(start N, target N) bool {
	checkVertex[N](mg, start)
	checkVertex[N](mg, target)
	visited := map[N]interface{}{start: nil}
	return mg.dfsLoop(start, target, visited)
}
