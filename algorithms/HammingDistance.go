package algorithms

import (
	"strings"
)

func flip(bit byte) string {
	if bit == '0' {
		return "1"
	}

	return "0"
}

func HammingDistance(bits string, uf UnionFind[string]) {
	bits = strings.ReplaceAll(bits, " ", "")
	l := len(bits)
	var temp string

	if _, exist := uf.nodes[bits]; exist {
		return
	}

	uf.nodes[bits] = bits
	uf.leaders[bits] = []string{bits}

	// 1 flip
	for i := range bits {
		temp = bits[:i] + flip(bits[i]) + bits[i+1:]

		if _, exist := uf.nodes[temp]; exist && !uf.Find(bits, temp) {
			uf.Union(bits, temp)
		}
	}

	// 2 flips
	for j := 0; j < l-1; j++ {
		for t := j + 1; t < l; t++ {
			temp = bits[:j] + flip(bits[j]) + bits[j+1:t] + flip(bits[t]) + bits[t+1:]

			if _, exist := uf.nodes[temp]; exist && !uf.Find(bits, temp) {
				uf.Union(bits, temp)
			}
		}
	}
}
