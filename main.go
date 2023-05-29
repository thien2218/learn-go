package main

import (
	"fmt"

	"github.com/thien2218/learn-go/algorithms"
)

func main() {
	arr := []int{9, 11, 14, 8, 19, 4, 2, 17, 16, 7, 12, 6, 5, 15, 13, 3, 1, 20, 18, 10}

	num := algorithms.QuickSelect(arr, 0, len(arr)-1, 19)
	fmt.Println(num)
}
