package algorithms

type heap []int

func NewHeap(arr []int) heap {
	var h heap = arr
	h.heapify()
	return h
}

func (h heap) heapify() {}

func (h heap) Insert() {}

func (h heap) Delete() {}

func HeapSort(arr []int) {}
