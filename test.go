package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/thien2218/learn-go/algorithms"
)

func test() {
	file, err := os.Open("excercises/twosum.txt")
	handleError(err)
	defer file.Close()

	arr := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		handleError(err)
		arr = append(arr, num)
	}

	count := algorithms.TwoSumInterval(arr)
	fmt.Println(count)
}
