package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/thien2218/learn-go/algorithms"
)

func main() {
	file, err := os.Open("excercises/test.txt")
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	capacity, err := strconv.Atoi(strings.Fields(scanner.Text())[0])
	handleError(err)

	items := make([][2]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		value, err := strconv.Atoi(fields[0])
		handleError(err)

		weight, err := strconv.Atoi(fields[1])
		handleError(err)

		item := [2]int{value, weight}
		items = append(items, item)
	}

	fmt.Println(algorithms.Knapsack(items, capacity))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
