package algorithms

func merge(left []int, right []int) []int {
	nLeft := len(left)
	nRight := len(right)

	arr := make([]int, nLeft+nRight)

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

func divide(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return append(make([]int, 1), arr[0])
	}

	mid := n / 2

	return merge(divide(arr[:mid]), divide(arr[mid:]))
}
