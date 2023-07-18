package algorithms

import "math/rand"

func partition(arr []int, left, right int) int {
	i := left + 1

	for j := left + 1; j <= right; j++ {
		if arr[j] < arr[left] {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i-1], arr[left] = arr[left], arr[i-1]
	return i - 1
}

func QuickSort(arr []int, left, right int) {
	if right-left <= 0 {
		return
	}

	pivot := rand.Intn(right-left+1) + left
	arr[left], arr[pivot] = arr[pivot], arr[left]

	i := partition(arr, left, right)

	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
