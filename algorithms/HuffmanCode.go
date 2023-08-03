package algorithms

import "sort"

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	Value  string
	Weight int
}

func pop(nodes *[]TreeNode) TreeNode {
	first := (*nodes)[0]
	(*nodes) = (*nodes)[1:]
	return first
}

func minSum(arr1, arr2 *[]TreeNode) (*TreeNode, *TreeNode) {
	mins := [2]TreeNode{}

	for j := 0; j < 2; j++ {
		if len(*arr2) == 0 || len(*arr1) > 0 && (*arr1)[0].Weight < (*arr2)[0].Weight {
			mins[j] = pop(arr1)
		} else {
			mins[j] = pop(arr2)
		}
	}

	return &mins[0], &mins[1]
}

func HuffmanCode(nodes []TreeNode) *TreeNode {
	if len(nodes) < 2 {
		return nil
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].Weight < nodes[j].Weight
	})

	sums := make([]TreeNode, 0)
	n := len(nodes)

	for i := 0; i < n-1; i++ {
		min1, min2 := minSum(&nodes, &sums)
		sum := new(TreeNode)

		sum.left = min1
		sum.right = min2
		sum.Weight = min1.Weight + min2.Weight

		sums = append(sums, *sum)
	}

	return &sums[0]
}
