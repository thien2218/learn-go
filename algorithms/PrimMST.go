package algorithms

type primMinHeap[V Node] struct {
	heap []Edge[V]
	pos  map[V]int
}

// PRIVATE
func (mh *primMinHeap[V]) swap(id1, id2 int) {
	mh.heap[id1], mh.heap[id2] = mh.heap[id2], mh.heap[id1]
	mh.pos[mh.heap[id1].EndVertex] = id1
	mh.pos[mh.heap[id2].EndVertex] = id2
}

func (mh *primMinHeap[V]) moveUp(id int) {
	heap := mh.heap

	for pid := (id - 1) / 2; pid >= 0 && heap[id].Weight < heap[pid].Weight; pid = (id - 1) / 2 {
		mh.swap(id, pid)
		id = pid
	}
}

func (mh *primMinHeap[V]) moveDown(id int) {
	var left, right int
	heap := mh.heap
	l := len(heap)

	for id < l && (left < l || right < l) {
		min := id
		left = 2*id + 1
		right = 2*id + 2

		if left < l && heap[left].Weight < heap[min].Weight {
			min = left
		}
		if right < l && heap[right].Weight < heap[min].Weight {
			min = right
		}
		if min != id {
			mh.swap(id, min)
			id = min
		} else {
			break
		}
	}
}

// PUBLIC
func (mh *primMinHeap[V]) Insert(edge Edge[V]) {
	id, exist := mh.pos[edge.EndVertex]

	if exist && mh.heap[id].Weight > edge.Weight {
		mh.heap[id] = edge
		mh.moveUp(id)
	} else if !exist {
		mh.heap = append(mh.heap, edge)
		last := len(mh.heap) - 1
		mh.pos[edge.EndVertex] = last
		mh.moveUp(last)
	}
}

func (mh *primMinHeap[V]) Extract() Edge[V] {
	l := len(mh.heap)
	first := mh.heap[0]

	mh.swap(0, l-1)
	mh.heap = mh.heap[:l-1]
	delete(mh.pos, first.EndVertex)
	mh.moveDown(0)

	return first
}

func PrimMst[V Node](graph Graph[V], source V) int {
	minHeap := new(primMinHeap[V])
	minHeap.heap = []Edge[V]{{EndVertex: source, Weight: 0}}
	minHeap.pos = map[V]int{source: 0}

	tracker := map[V]bool{source: true}
	minCost := 0

	for {
		minEdge := minHeap.Extract()
		tracker[minEdge.EndVertex] = true
		minCost += int(minEdge.Weight)

		if len(graph.vertices) == len(tracker) {
			break
		}

		for _, edge := range graph.GetEdges(minEdge.EndVertex) {
			if _, tracked := tracker[edge.EndVertex]; tracked {
				continue
			}

			minHeap.Insert(edge)
		}
	}

	return minCost
}
