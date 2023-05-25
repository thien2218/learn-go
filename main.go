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
	arr := make([]int, 0, 100000)
	file, err := os.Open("intArr.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		arr = append(arr, num)
	}

	_, count := algorithms.Inversion(arr)
	fmt.Println(count)
}
