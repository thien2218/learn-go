package algorithms

func TwoSumInterval(nums []int) int {
	record := make(map[int][]int)
	sums := make(map[int]bool)

	for _, num := range nums {
		for i := -1; i <= 1; i++ {
			oppositeKey := ((0 - num) / 10000) + i

			if _, exist := record[oppositeKey]; exist {
				for _, value := range record[oppositeKey] {
					oppositeNum := oppositeKey*10000 + value
					sum := num + oppositeNum

					if _, exist = sums[sum]; sum <= 10000 && sum >= -10000 && !exist && oppositeNum != num {
						sums[sum] = true
					}
				}
			}

			key := num / 10000
			value := num % 10000

			record[key] = append(record[key], value)
		}
	}

	return len(sums)
}
