package store

import (
	"errors"

	"itib/lab9/cluster"
)

var areas []*cluster.Area

func Clear(id int) {
	areas[id] = nil
}

func AddArea() (int, error) {
	areas = append(areas, cluster.CreateArea())
	return len(areas) - 1, nil
}

func GetArea(id int) (*cluster.Area, error) {
	if id < 0 || id >= len(areas) {
		return nil, errors.New("incorrect area id")
	}

	return areas[id], nil
}

func GetDistanceFunction(id int) (cluster.Metric, error) {
	switch id {
	case 1:
		return cluster.EuclideanDistance, nil
	case 2:
		return cluster.ManhattanDistance, nil
	case 3:
		return cluster.ChebyshevDistance, nil
	default:
		return nil, errors.New("incorrect distance func id")
	}
}
