package algorithms

import "math"

type PathWVertex[V Node] struct {
	Weight float64
	Edges  [2]Edge[V]
}

func Mwis[V Node](pathGraph map[V]PathWVertex[V], start V) []float64 {
	cache := []float64{0, pathGraph[start].Weight}
	vertex := pathGraph[start]

	n := len(pathGraph)
	for i := 1; i < n; i++ {
		next := vertex.Edges[1].EndVertex
		vertex = pathGraph[next]

		cachedItem := math.Max(cache[i], cache[i-1]+vertex.Weight)
		cache = append(cache, cachedItem)
	}

	return cache
}

func MwisReconstruct[V Node](pathGraph map[V]PathWVertex[V], end V, cache []float64) []int {
	vertexIds := make([]int, 0)
	i := len(cache) - 1
	vertex := pathGraph[end]

	for i >= 1 {
		if i > 1 && cache[i-1] >= cache[i-2]+vertex.Weight {
			vertex = pathGraph[vertex.Edges[0].EndVertex]
			i--
		} else {
			vertexIds = append(vertexIds, i)
			vertex = pathGraph[vertex.Edges[0].EndVertex]
			vertex = pathGraph[vertex.Edges[0].EndVertex]
			i -= 2
		}
	}

	return vertexIds
}
