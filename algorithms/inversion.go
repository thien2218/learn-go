package algorithms

func countAndMerge(left []int, right []int) ([]int, int) {
	nLeft := len(left)
	nRight := len(right)
	count := 0

	arr := make([]int, 0, nLeft+nRight)

	var i int
	var j int

	for {
		if left[i] <= right[j] {
			arr = append(arr, left[i])
			i++
		} else {
			arr = append(arr, right[j])
			count += nLeft - i
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

	return arr, count
}

func Inversion(arr []int) ([]int, int) {
	n := len(arr)

	if n == 1 {
		return []int{arr[0]}, 0
	}

	mid := n / 2

	left, c1 := Inversion(arr[:mid])
	right, c2 := Inversion(arr[mid:])
	arr, c3 := countAndMerge(left, right)
	return arr, c1 + c2 + c3
}
