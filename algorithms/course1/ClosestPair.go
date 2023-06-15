package algorithms

import (
	"math"
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

func closestSplitPair(p Point, py []Point, min float32) ([2]Point, float32) {
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
	// Check for base case when length = 2 (right - left = 1)
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

func FindClosestPair(points []Point) {

}
