package algorithms

import "log"

// CONSTRUCTOR
func NewHeap[N Node](heapType string, arr ...N) Heap[N] {
	if heapType != "min" && heapType != "max" {
		log.Fatal("Invalid type of heap.")
	}

	var h Heap[N]
	h.nodes = arr
	h.heapType = heapType
	h.heapify()
	return h
}

// PRIVATE
func (h Heap[N]) heapify() {
	nodes := h.nodes
	l := len(nodes)

	for i := (l / 2) - 1; l > 0 && i >= 0; i-- {
		h.moveDown(i)
	}
}

func (h Heap[N]) comp(val1, val2 N) bool {
	if h.heapType == "min" {
		return val1 < val2
	}

	return val1 > val2
}

func (h *Heap[N]) moveUp(id int) {
	nodes := (*h).nodes

	for pid := (id - 1) / 2; pid >= 0 && (*h).comp(nodes[id], nodes[pid]); pid = (id - 1) / 2 {
		nodes[id], nodes[pid] = nodes[pid], nodes[id]
		id = pid
	}
}

func (h *Heap[N]) moveDown(id int) {
	var left, right int
	nodes := (*h).nodes
	l := len(nodes)

	for left < l || right < l {
		minOrMax := id
		left = 2*id + 1
		right = 2*id + 2

		if left < l && (*h).comp(nodes[left], nodes[minOrMax]) {
			minOrMax = left
		}
		if right < l && (*h).comp(nodes[right], nodes[minOrMax]) {
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
func (h *Heap[N]) Insert(node N) {
	(*h).nodes = append((*h).nodes, node)
	(*h).moveUp(len((*h).nodes) - 1)
}

func (h *Heap[N]) Update(node N, newValue N) bool {
	for i, heapNode := range (*h).nodes {
		if heapNode == node {
			(*h).nodes[i] = newValue
			(*h).heapify()
			return true
		}
	}

	return false
}

func (h *Heap[N]) Extract() N {
	l := len((*h).nodes)
	first := (*h).nodes[0]

	(*h).nodes[0], (*h).nodes[l-1] = (*h).nodes[l-1], (*h).nodes[0]
	(*h).nodes = (*h).nodes[:l-1]
	(*h).moveDown(0)

	return first
}
