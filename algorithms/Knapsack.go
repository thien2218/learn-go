package algorithms

import "sort"

type LinkedList[T interface{}] struct {
	item T
	prev *LinkedList[T]
	next *LinkedList[T]
}

func (node1 *LinkedList[T]) Insert(node2 *LinkedList[T]) {
	if node2 == nil {
		node1.next = node2
		return
	}

	node2.next = node1.next
	node2.prev = node1

	node1.next = node2

	if node2.next != nil {
		node2.next.prev = node2
	}
}

func (node1 *LinkedList[T]) Connect(node2 *LinkedList[T]) {
	node1.next = node2

	if node2 != nil {
		node2.prev = node1
	}
}

func (node *LinkedList[T]) Disconnect() *LinkedList[T] {
	next := node.next
	node.next = nil
	return next
}

func Knapsack(items [][2]int, capacity int) [2]int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] < items[j][1]
	})

	results := new(LinkedList[[2]int])
	results.item = items[0]

	curr := results
	n := len(items)

	for i := 1; i < n; i++ {
		item := items[i]
		accs := [][2]int{item}

		temp := results
		for temp != nil && temp.item[1]+item[1] <= capacity {
			newItem := [2]int{temp.item[0] + item[0], temp.item[1] + item[1]}
			accs = append(accs, newItem)
			temp = temp.next
		}

		curr = insertOrContract(curr, accs[0])
		temp = curr

		for _, acc := range accs[1:] {
			temp = insertOrContract(temp, acc)
		}
	}

	for curr.next != nil {
		curr = curr.next
	}

	return curr.item
}

func insertOrContract(curr *LinkedList[[2]int], acc [2]int) *LinkedList[[2]int] {
	for curr.next != nil && curr.next.item[1] < acc[1] {
		curr = curr.next
	}

	newNode := new(LinkedList[[2]int])
	newNode.item = acc

	if acc[0] > curr.item[0] && (curr.next == nil || acc[0] < curr.next.item[0] && acc[1] < curr.next.item[1]) {
		curr.Insert(newNode)
	} else {
		contract(curr, newNode)
	}

	return curr
}

func contract(node, newNode *LinkedList[[2]int]) {
	if node.next != nil && node.next.item[0] <= newNode.item[0] {
		disconn := node.Disconnect()

		for disconn != nil && disconn.item[0] <= newNode.item[0] {
			next := disconn.next
			disconn = nil
			disconn = next
		}

		newNode.Connect(disconn)
		node.Connect(newNode)
	}
}
