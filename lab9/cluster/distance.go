package cluster

import (
	"math"
)

type Metric = func(Point, Point) float64

func EuclideanDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

func ChebyshevDistance(p1 Point, p2 Point) float64 {
	return math.Max(math.Abs(p1.X-p2.X), math.Abs(p1.Y-p2.Y))
}

func ManhattanDistance(p1 Point, p2 Point) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p2.Y-p2.Y)
}
