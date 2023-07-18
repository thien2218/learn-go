package algorithms

func merge(left []int, right []int) []int {
	nLeft := len(left)
	nRight := len(right)

	arr := make([]int, 0, nLeft+nRight)

	var i int
	var j int

	for {
		if left[i] <= right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			j++
		}

		if i == nLeft {
			arr = append(arr, right[j:]...)
			break
		} else if j == nRight {
			arr = append(arr, left[i:]...)
			break
		}
	}

	return arr
}

func MergeSort(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return arr
	}

	mid := n / 2

	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
