package algorithms

import "log"

type heap[N Node] struct {
	nodes    []N
	heapType string
}

// CONSTRUCTOR
func NewHeap[N Node](heapType string, arr ...N) heap[N] {
	if heapType != "min" && heapType != "max" {
		log.Fatal("Invalid type of heap.")
	}

	var h heap[N]
	h.nodes = arr
	h.heapType = heapType
	h.heapify()
	return h
}

// PRIVATE
func (h heap[N]) heapify() {
	nodes := h.nodes
	l := len(nodes)

	for i := (l / 2) - 1; l > 0 && i >= 0; i-- {
		h.moveDown(i)
	}
}

func (h heap[N]) comp(val1, val2 N) bool {
	if h.heapType == "min" {
		return val1 < val2
	}

	return val1 > val2
}

func (h *heap[N]) moveUp(id int) {
	nodes := (*h).nodes

	for pid := (id - 1) / 2; pid >= 0 && (*h).comp(nodes[id], nodes[pid]); pid = (id - 1) / 2 {
		nodes[id], nodes[pid] = nodes[pid], nodes[id]
		id = pid
	}
}

func (h *heap[N]) moveDown(id int) {
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
func (h *heap[N]) Insert(node N) {
	(*h).nodes = append((*h).nodes, node)
	(*h).moveUp(len((*h).nodes) - 1)
}

func (h *heap[N]) Update(node N, newValue N) bool {
	for i, heapNode := range (*h).nodes {
		if heapNode == node {
			(*h).nodes[i] = newValue
			(*h).heapify()
			return true
		}
	}

	return false
}

func (h *heap[N]) Extract() N {
	l := len((*h).nodes)
	first := (*h).nodes[0]

	(*h).nodes[0], (*h).nodes[l-1] = (*h).nodes[l-1], (*h).nodes[0]
	(*h).nodes = (*h).nodes[:l-1]
	(*h).moveDown(0)

	return first
}
