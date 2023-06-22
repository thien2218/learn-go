package algorithms

import (
	"math"
	"unsafe"
)

type Point struct {
	x float32
	y float32
}

func (p1 Point) distanceTo(p2 Point) float32 {
	xDiff := math.Abs(float64(p2.x*p2.x - p1.x*p1.x))
	yDiff := math.Abs(float64(p2.y*p2.y - p1.y*p1.y))

	euclidean := float32(math.Sqrt(xDiff + yDiff))
	return euclidean
}

const (
	signMask uint32 = 1 << 31
	fullMask uint32 = math.MaxUint32
)

// float32ToUint32 converts a float32 number to a uint32 number
func float32ToUint32(f float32) uint32 {
	u := *(*uint32)(unsafe.Pointer(&f))
	if u&signMask == signMask {
		u = ^u
	} else {
		u |= signMask
	}
	return u
}

// getBit returns the bit at the given position of the uint32 number
func getBit(u uint32, pos int) int {
	return int((u >> pos) & 1)
}

// radixSortX sorts the given slice of Point structs based on their x values using Radix Sort
func radixSortX(arr []Point) []Point {
	n := len(arr)
	output := make([]Point, n)

	for pos := 0; pos < 32; pos++ {
		count := [2]int{0, 0}

		for _, p := range arr {
			u := float32ToUint32(p.x)
			count[getBit(u, pos)]++
		}

		count[1] += count[0]

		for i := n - 1; i >= 0; i-- {
			u := float32ToUint32(arr[i].x)
			bit := getBit(u, pos)
			count[bit]--
			output[count[bit]] = arr[i]
		}
	}

	return output
}

// radixSortY sorts the given slice of Point structs based on their y values using Radix Sort
func radixSortY(arr []Point) []Point {
	n := len(arr)
	output := make([]Point, n)

	for pos := 0; pos < 32; pos++ {
		count := [2]int{0, 0}

		for _, p := range arr {
			u := float32ToUint32(p.y)
			count[getBit(u, pos)]++
		}

		count[1] += count[0]

		for i := n - 1; i >= 0; i-- {
			u := float32ToUint32(arr[i].y)
			bit := getBit(u, pos)
			count[bit]--
			output[count[bit]] = arr[i]
		}
	}

	return output
}

func closestSplitPair(p Point, py []Point, min float32) ([2]Point, float32) {
	// TODO: Use optimized sorting algorithm to refine closestSplitPair
	// 		subroutine running time

	sy := make([]Point, 0)
	var pair [2]Point

	for _, point := range py {
		if point.x < p.x+min && point.x > p.x-min {
			sy = append(sy, point)
		}
	}

	l := len(sy)

	for i := 0; i < l; i++ {
		for j := 1; j <= int(math.Min(7, float64(l-i-1))); j++ {
			d := sy[i].distanceTo(sy[i+j])
			if d < min {
				min = d
				pair[0], pair[1] = sy[i], sy[i+j]
			}
		}
	}

	return pair, min
}

func closestPair(px, py []Point) ([2]Point, float32) {
	l := len(px)

	if l == 2 {
		pair := [2]Point{px[0], px[1]}
		return pair, px[0].distanceTo(px[1])
	}

	mid := l / 2

	_, d1 := closestPair(px[:mid+1], py)
	_, d2 := closestPair(px[mid:], py)

	min := float32(math.Min(float64(d1), float64(d2)))
	return closestSplitPair(px[mid+1], py, min)
}

func FindClosestPair(points []Point) (Point, Point, float32) {
	px := radixSortX(points)
	py := radixSortY(points)

	p, d := closestPair(px, py)

	return p[0], p[1], d
}
