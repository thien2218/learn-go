package algorithms

type minHeapGraph[V Node] []Edge[V]

// Constructor
func NewMinHeapGraph[V Node](edges []Edge[V]) minHeapGraph[V] {
	var heap minHeapGraph[V] = edges
	heap.heapify()
	return heap
}

// PRIVATE
func (h minHeapGraph[V]) heapify() {
	l := len(h)

	for i := (l / 2) - 1; l > 0 && i >= 0; i-- {
		h.moveDown(i)
	}
}

func (h minHeapGraph[V]) moveUp(id int) {
	for pid := (id - 1) / 2; pid >= 0 && h[id].Weight < h[pid].Weight; pid = (id - 1) / 2 {
		h[id], h[pid] = h[pid], h[id]
		id = pid
	}
}

func (h minHeapGraph[V]) moveDown(id int) {
	var left, right int
	l := len(h)

	for left < l || right < l {
		min := id
		left = 2*id + 1
		right = 2*id + 2

		if left < l && h[left].Weight < h[min].Weight {
			min = left
		}
		if right < l && h[right].Weight < h[min].Weight {
			min = right
		}
		if min != id {
			h[min], h[id] = h[id], h[min]
			id = min
		} else {
			break
		}
	}
}

// PUBLIC
func (h *minHeapGraph[V]) Insert(edge Edge[V]) {
	*h = append(*h, edge)
	h.moveUp(len(*h) - 1)
}

func (h *minHeapGraph[V]) Extract() Edge[V] {
	l := len(*h)
	first := (*h)[0]

	(*h)[0], (*h)[l-1] = (*h)[l-1], (*h)[0]
	*h = (*h)[:l-1]
	h.moveDown(0)

	return first
}

func Dijkstra[V Node](graph Graph[V], source V, target V) float64 {
	// Initialize visited tracker and a queue just like BFS
	// but instead of using a queue, we use a heap
	heap := NewMinHeapGraph[V](graph.GetEdges(source))
	// and a map of distances as the tracker to keep track of the shortest
	// distances from the source to other vertices
	distances := map[V]float64{source: 0}

	var minEdge Edge[V]

	// Loop until the heap is empty
	for len(heap) > 0 && minEdge.EndVertex != target {
		minEdge = heap.Extract()
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
