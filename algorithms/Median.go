package algorithms

import "fmt"

func Median(arr []int) {
	var median, sum int
	minHeap := NewOrderedMinHeap[int]() // used to get minimum of the bigger half
	maxHeap := NewOrderedMaxHeap[int]() // used to get maximum of the smaller half

	for _, num := range arr {
		lMin := len(minHeap.nodes)
		lMax := len(maxHeap.nodes)

		if median == 0 {
			median = num
		} else if num > median {
			if lMin == lMax+1 {
				minHeap.Insert(num)
				maxHeap.Insert(median)
				median = minHeap.Extract()
			} else {
				minHeap.Insert(num)
			}
		} else if num < median {
			if lMax == lMin {
				maxHeap.Insert(num)
				minHeap.Insert(median)
				median = maxHeap.Extract()
			} else {
				maxHeap.Insert(num)
			}
		}

		sum += median
	}

	fmt.Println(sum)
}
