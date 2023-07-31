package algorithms

import (
	"fmt"
	"sort"
)

type Link[V Node] struct {
	vertex V
	edge   Edge[V]
}

func NewLink[V Node](vertex1, vertex2 V, weight float64) *Link[V] {
	edge := new(Edge[V])
	edge.EndVertex = vertex2
	edge.Weight = weight

	link := new(Link[V])
	link.vertex = vertex1
	link.edge = *edge

	return link
}

func KClustering[V Node](links []Link[V], uf UnionFind[V], k int) float64 {
	distances := make(map[string]float64)

	sort.SliceStable(links, func(i, j int) bool {
		return links[i].edge.Weight < links[j].edge.Weight
	})

	for _, link := range links {
		v1, v2 := link.vertex, link.edge.EndVertex
		cost := link.edge.Weight

		if !uf.Find(v1, v2) {
			if len(uf.leaders) > k {
				uf.Union(v1, v2)
			} else {
				str1 := fmt.Sprintf("%v", uf.nodes[v1])
				str2 := fmt.Sprintf("%v", uf.nodes[v2])

				key1 := str1 + "/" + str2
				key2 := str2 + "/" + str1

				_, exist1 := distances[key1]
				_, exist2 := distances[key2]

				if !exist1 && !exist2 {
					distances[key1] = cost
				}
			}
		}
	}

	var max float64
	for _, distance := range distances {
		if distance > max {
			max = distance
		}
	}

	return max
}
