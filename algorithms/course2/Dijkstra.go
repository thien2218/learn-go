package algorithms

import "math"

type minHeapGraph[V Vertex] []Edge[V]

// Constructor
func NewMinHeapGraph[V Vertex](edges []Edge[V]) minHeapGraph[V] {
	var heap minHeapGraph[V] = edges
	heap.heapify()
	return heap
}

// PRIVATE
func (h minHeapGraph[V]) heapify() {
	i := 0
	l := len(h)

	left := 2*i + 1
	right := 2*i + 2

	for l > 0 && (left < l || right < l) {
		if left < l && right < l {
			if h[left].Weight < h[i].Weight && h[left].Weight < h[right].Weight {
				h.moveUp(left)
			} else if h[right].Weight < h[i].Weight {
				h.moveUp(right)
			}
		} else if left < l && h[left].Weight < h[i].Weight {
			h.moveUp(left)
		} else if right < l && h[right].Weight < h[i].Weight {
			h.moveUp(right)
		} else {
			break
		}

		i++
		left = 2*i + 1
		right = 2*i + 2
	}
}

func (h minHeapGraph[V]) moveUp(id int) {
	parentId := int(math.Floor(float64(id-1) / 2))

	for parentId >= 0 && h[id].Weight < h[parentId].Weight {
		h[id], h[parentId] = h[parentId], h[id]
		id = parentId
		parentId = int(math.Floor(float64(id-1) / 2))
	}
}

func (h minHeapGraph[V]) moveDown(id int) {
	l := len(h)
	left := 2*id + 1
	right := 2*id + 2

	for left < l || right < l {
		if left < l && right < l {
			if h[left].Weight < h[id].Weight && h[left].Weight < h[right].Weight {
				h[left], h[id] = h[id], h[left]
				id = left
			} else if h[right].Weight < h[id].Weight {
				h[right], h[id] = h[id], h[right]
				id = right
			} else {
				break
			}
		} else if left < l && h[left].Weight < h[id].Weight {
			h[left], h[id] = h[id], h[left]
			id = left
		} else if right < l && h[right].Weight < h[id].Weight {
			h[right], h[id] = h[id], h[right]
			id = right
		} else {
			break
		}

		left = 2*id + 1
		right = 2*id + 2
	}
}

// PUBLIC
func (h *minHeapGraph[V]) Insert(edge Edge[V]) {
	*h = append(*h, edge)
	h.moveUp(len(*h) - 1)
}

func (h *minHeapGraph[V]) Pop() Edge[V] {
	l := len(*h)
	first := (*h)[0]

	(*h)[0], (*h)[l-1] = (*h)[l-1], (*h)[0]
	*h = (*h)[:l-1]
	h.moveDown(0)

	return first
}

func Dijkstra[V Vertex](graph IGraph[V], source V, target V) float64 {
	// Initialize visited tracker and a queue just like BFS
	// but instead of using a queue, we use a heap
	heap := NewMinHeapGraph[V](graph.GetEdges(source))
	// and a map of distances as the tracker to keep track of the shortest
	// distances from the source to other vertices
	distances := map[V]float64{source: 0}

	var minEdge Edge[V]

	// Loop until the heap is empty
	for len(heap) > 0 && minEdge.EndVertex != target {
		minEdge = heap.Pop()
		distances[minEdge.EndVertex] = minEdge.Weight

		for _, edge := range graph.GetEdges(minEdge.EndVertex) {
			if _, exist := distances[edge.EndVertex]; !exist {
				edge.Weight = edge.Weight + minEdge.Weight
				heap.Insert(edge)
			}
		}
	}

	if _, exist := distances[target]; exist {
		return distances[target]
	}

	return 1000000
}
