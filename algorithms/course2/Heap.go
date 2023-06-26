package algorithms

import (
	"log"
	"math"
)

type heap struct {
	nodes    []int
	heapType string
}

// CONSTRUCTOR
func NewHeap(arr []int, heapType string) heap {
	if heapType != "min" && heapType != "max" {
		log.Fatal("Invalid type of heap.")
	}

	var h heap
	h.nodes = arr
	h.heapType = heapType
	h.heapify()
	return h
}

// PRIVATE
func (h heap) heapify() {
	i := 0
	nodes := h.nodes
	l := len(nodes)

	left := 2*i + 1
	right := 2*i + 2

	for l > 0 && (left < l || right < l) {
		if left < l && right < l {
			if h.comp(nodes[left], nodes[i]) && h.comp(nodes[left], nodes[right]) {
				h.moveUp(left)
			} else if h.comp(nodes[right], nodes[i]) {
				h.moveUp(right)
			}
		} else if left < l && h.comp(nodes[left], nodes[i]) {
			h.moveUp(left)
		} else if right < l && h.comp(nodes[right], nodes[i]) {
			h.moveUp(right)
		} else {
			break
		}

		i++
		left = 2*i + 1
		right = 2*i + 2
	}
}

func (h heap) comp(val1, val2 int) bool {
	if h.heapType == "min" {
		return val1 < val2
	}

	return val1 > val2
}

func (h heap) moveUp(id int) {
	nodes := h.nodes

	for parentId := int(math.Floor(float64(id-1) / 2)); h.comp(nodes[id], nodes[parentId]); {
		nodes[id], nodes[parentId] = nodes[parentId], nodes[id]
		id = parentId
		parentId = int(math.Floor(float64(id-1) / 2))
	}
}

func (h heap) moveDown(id int) {
	nodes := h.nodes
	l := len(nodes)

	left := 2*id + 1
	right := 2*id + 2

	for left < l || right < l {
		if left < l && right < l {
			if !h.comp(nodes[left], nodes[id]) && !h.comp(nodes[left], nodes[right]) {
				nodes[left], nodes[id] = nodes[id], nodes[left]
				id = left
			} else if !h.comp(nodes[right], nodes[id]) {
				nodes[right], nodes[id] = nodes[id], nodes[right]
				id = right
			}
		} else if left < l && !h.comp(nodes[left], nodes[id]) {
			nodes[left], nodes[id] = nodes[id], nodes[left]
			id = left
		} else if right < l && !h.comp(nodes[right], nodes[id]) {
			nodes[right], nodes[id] = nodes[id], nodes[right]
			id = right
		} else {
			break
		}

		left = 2*id + 1
		right = 2*id + 2
	}
}

// PUBLIC
func (h heap) Insert(node int) {
	h.nodes = append(h.nodes, node)
	h.moveUp(len(h.nodes) - 1)
}

func (h heap) Update(node int, newValue int) bool {
	for i, heapNode := range h.nodes {
		if heapNode == node {
			h.nodes[i] = newValue
			h.heapify()
			return true
		}
	}

	return false
}

func (h heap) DeleteFirst() int {
	l := len(h.nodes)
	ptr := &h
	first := h.nodes[0]

	h.nodes[0], h.nodes[l-1] = h.nodes[l-1], h.nodes[0]
	h.nodes = ptr.nodes[:l-1]
	h.moveDown(0)

	return first
}
