package cluster

import (
	"itib/lab9/utils"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Cluster struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Points []Point `json:"points"`
}

type Area struct {
	points   []Point
	clusters []Point
	age      uint64
}

type average struct {
	sum   Point
	count uint64
}

func CreateArea() *Area {
	return &Area{
		points:   make([]Point, 0),
		clusters: make([]Point, 0),
	}
}

func (a *Area) AddPoint(point Point) {
	a.points = append(a.points, point)
}

func (a *Area) AddCluster(cluster Point) {
	a.clusters = append(a.clusters, cluster)
}

func (a *Area) DoStep(distanceFunction Metric) bool {
	if len(a.points) == 0 || len(a.clusters) == 0 {
		return true
	}

	a.age += 1
	newClusters := make([]average, len(a.clusters))

	for _, point := range a.points {
		minDistance := -42.
		minCluster := -42

		for c, cluster := range a.clusters {
			distance := distanceFunction(point, cluster)

			if minCluster == -42 || distance < minDistance {
				minDistance = distance
				minCluster = c
			}
		}

		newClusters[minCluster].sum.X += point.X
		newClusters[minCluster].sum.Y += point.Y
		newClusters[minCluster].count += 1
	}

	isFinished := true

	for i := range newClusters {
		prevX := a.clusters[i].X
		prevY := a.clusters[i].Y

		if newClusters[i].count != 0 {
			a.clusters[i].X = newClusters[i].sum.X / float64(newClusters[i].count)
			a.clusters[i].Y = newClusters[i].sum.Y / float64(newClusters[i].count)
		}

		if (utils.Round(a.clusters[i].X-prevX, 100) != 0) || (utils.Round(a.clusters[i].Y-prevY, 100) != 0) {
			isFinished = false
		}
	}

	return isFinished
}

func (a *Area) Learn(distanceFunction Metric, maxIterations uint64) bool {
	for a.age = 0; a.age < maxIterations; {
		if a.DoStep(distanceFunction) {
			return true
		}
	}
	return false
}

func (a *Area) Clear() {
	a.points = a.points[:0]
	a.clusters = a.clusters[:0]
}

func (a *Area) GetClusterPoints() []Point {
	return a.clusters
}

func (a *Area) GetClusters(distanceFunction Metric) []Cluster {
	if len(a.clusters) == 0 {
		return nil
	}

	clusters := make([]Cluster, len(a.clusters))

	for i := range a.clusters {
		clusters[i].X = a.clusters[i].X
		clusters[i].Y = a.clusters[i].Y
		clusters[i].Points = make([]Point, 0)
	}

	for i, point := range a.points {
		minDistance := -42.
		minCluster := -42

		for i, cluster := range a.clusters {
			distance := distanceFunction(point, cluster)
			if minCluster == -42 || distance < minDistance {
				minDistance = distance
				minCluster = i
			}
		}

		clusters[minCluster].Points = append(clusters[minCluster].Points, a.points[i])
	}

	return clusters
}
