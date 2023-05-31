package main

import (
	"bufio"
	// "fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "github.com/thien2218/learn-go/algorithms"
)

func main() {

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readGraphFromFile(graph map[int][]int, filename string) {
	file, err := os.Open(filename)
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		words := strings.Fields(line)

		for j, word := range words {
			n, err := strconv.Atoi(word)
			handleError(err)

			if j == 0 {
				graph[i+1] = make([]int, 0, 200)
			} else {
				graph[i+1] = append(graph[i+1], n)
			}
		}
	}

	handleError(scanner.Err())
}

// func readNumsFromFile(arr []int, filename string) {
// 	file, err := os.Open(filename)
// 	handleError(err)
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for i := 0; scanner.Scan(); i++ {
// 		txt := scanner.Text()
// 		num, err := strconv.Atoi(txt)
// 		handleError(err)
// 		arr = append(arr, num)
// 	}

// 	handleError(scanner.Err())
// }
