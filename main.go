package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/thien2218/learn-go/algorithms"
)

func main() {
	file, err := os.Open("excercises/mwis.txt")
	handleError(err)
	defer file.Close()

	var i int
	pathGraph := make(map[int]algorithms.PathWVertex[int])

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	for scanner.Scan() {
		i++

		line := scanner.Text()
		weight, err := strconv.ParseFloat(line, 64)
		handleError(err)

		vertex := new(algorithms.PathWVertex[int])
		vertex.Weight = weight
		vertex.Edges[0].EndVertex = i - 1
		vertex.Edges[1].EndVertex = i + 1

		pathGraph[i] = *vertex
	}

	cache := algorithms.Mwis[int](pathGraph, 1)
	vertexIds := algorithms.MwisReconstruct[int](pathGraph, len(pathGraph), cache)
	fmt.Println(vertexIds)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
