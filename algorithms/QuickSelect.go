package algorithms

import (
	"log"
	"math/rand"
)

func QuickSelect(arr []int, left, right, ith int) int {
	if ith < left+1 || ith > right+1 {
		log.Fatal("Invalid order statistic 'i'.")
	}

	pivot := rand.Intn(right-left+1) + left
	arr[left], arr[pivot] = arr[pivot], arr[left]

	i := partition(arr, left, right)

	switch {
	case i == ith-1:
		return arr[i]
	case i > ith-1:
		return QuickSelect(arr, left, i-1, ith)
	default:
		return QuickSelect(arr, i+1, right, ith)
	}
}
