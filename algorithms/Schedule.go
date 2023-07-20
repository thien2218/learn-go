package algorithms

import "sort"

func NonOptimalScheduling(jobs [][2]int) int {
	sort.SliceStable(jobs, func(i, j int) bool {
		diff1 := jobs[i][0] - jobs[i][1]
		diff2 := jobs[j][0] - jobs[j][1]

		if diff1 > diff2 || (diff1 == diff2 && jobs[i][0] > jobs[j][0]) {
			return true
		}

		return false
	})

	var acc int
	var total int

	for _, job := range jobs {
		acc += job[1]
		total += job[0] * acc
	}

	return total
}

func OptimalScheduling(jobs [][2]int) int {
	sort.SliceStable(jobs, func(i, j int) bool {
		ratio1 := float64(jobs[i][0]) / float64(jobs[i][1])
		ratio2 := float64(jobs[j][0]) / float64(jobs[j][1])

		return ratio1 > ratio2
	})

	var acc int
	var total int

	for _, job := range jobs {
		acc += job[1]
		total += job[0] * acc
	}

	return total
}
