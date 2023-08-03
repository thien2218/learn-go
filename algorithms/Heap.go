package algorithms

type Heap[T interface{}] struct {
	nodes   []T
	compare func(i, j int) bool
}

// CONSTRUCTOR

func NewOrderedMinHeap[N Node](arr ...N) *Heap[N] {
	h := new(Heap[N])
	h.nodes = arr
	h.compare = func(i, j int) bool {
		return h.nodes[i] < h.nodes[j]
	}

	h.heapify()
	return h
}

func NewOrderedMaxHeap[N Node](arr ...N) *Heap[N] {
	h := new(Heap[N])
	h.nodes = arr
	h.compare = func(i, j int) bool {
		return h.nodes[i] > h.nodes[j]
	}

	h.heapify()
	return h
}

func NewHeap[T interface{}](compare func(i, j int) bool, arr ...T) Heap[T] {
	var h Heap[T]
	h.nodes = arr
	h.compare = compare
	h.heapify()
	return h
}

// PRIVATE

func (h *Heap[T]) heapify() {
	nodes := h.nodes
	l := len(nodes)

	for i := (l / 2) - 1; l > 0 && i >= 0; i-- {
		h.moveDown(i)
	}
}

func (h *Heap[T]) moveUp(id int) {
	nodes := h.nodes

	for pid := (id - 1) / 2; pid >= 0 && h.compare(id, pid); pid = (id - 1) / 2 {
		nodes[id], nodes[pid] = nodes[pid], nodes[id]
		id = pid
	}
}

func (h *Heap[T]) moveDown(id int) {
	var left, right int
	nodes := h.nodes
	l := len(nodes)

	for left < l || right < l {
		minOrMax := id
		left = 2*id + 1
		right = 2*id + 2

		if left < l && h.compare(left, minOrMax) {
			minOrMax = left
		}
		if right < l && h.compare(right, minOrMax) {
			minOrMax = right
		}
		if minOrMax != id {
			nodes[minOrMax], nodes[id] = nodes[id], nodes[minOrMax]
			id = minOrMax
		} else {
			break
		}
	}
}

// PUBLIC

func (h *Heap[T]) Insert(node T) {
	h.nodes = append(h.nodes, node)
	h.moveUp(len(h.nodes) - 1)
}

func (h *Heap[T]) Extract() T {
	l := len(h.nodes)
	first := h.nodes[0]

	h.nodes[0], h.nodes[l-1] = h.nodes[l-1], h.nodes[0]
	h.nodes = h.nodes[:l-1]
	h.moveDown(0)

	return first
}
