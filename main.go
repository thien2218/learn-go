package main

import (
	"fmt"

	"github.com/thien2218/learn-go/algorithms"
)

func main() {
	res := algorithms.Karatsuba("00123", "4567")
	fmt.Println(res)
}
