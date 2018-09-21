package main

import (
	"fmt"
	"math"
)

//Point represents data
type Point []float64

//Distance returns distance between 'other' and 'point'
func (point Point) Distance(other Point) float64 {
	val := 0.0
	for ii := range point {
		val += math.Pow(point[ii]-other[ii], 2)
	}
	return math.Sqrt(val)
}

func main() {

	points := []Point{
		Point{3, 3},
		Point{2, 3},
		Point{2, 2},
		Point{1, 1},
		Point{0, 0},
		Point{1, -1},
		Point{-1, -1},
		Point{-3, -2},
		Point{-2, 0},
		Point{-2, -2},
	}

	centroids := []Point{
		Point{2, 2},
		Point{1, 1},
		Point{-3, -2},
		Point{-2, -2},
	}

	centroids = kmeans(points, centroids, 10, 1)

	fmt.Println(centroids)
	fmt.Println(classify(points, centroids))

}

func kmeans(points []Point, centroids []Point, maxIter int, eta float64) []Point {
	centroidsMap := classify(points, centroids)
	updatedCentroids := updateCentroids(centroids, centroidsMap)
	if maxIter > 0 && !converged(centroids, updatedCentroids, eta) {
		return kmeans(points, updatedCentroids, maxIter-1, eta)
	} else {
		return updatedCentroids
	}
}

func converged(centroids []Point, updatedCentroids []Point, eta float64) bool {
	ch := make(chan bool, len(centroids))
	for ii := range centroids {
		go func(jj int, ch chan bool) {
			if centroids[jj].Distance(updatedCentroids[jj]) < eta {
				ch <- true
			} else {
				ch <- false
			}
		}(ii, ch)
	}

	for range centroids {
		if result := <-ch; !result {
			return false
		}
	}
	return true
}

//Result encapsulates a 'data point', 'nearest centroid index', and their distance
type Result struct {
	point       Point
	centroidInd int
	distance    float64
}

func classify(points []Point, centroids []Point) map[int][]Point {

	ch := make(chan Result)

	for ii := range points {
		go func(jj int, ch chan Result) {
			ch <- Result{
				centroidInd: findClosest(points[jj], centroids),
				point:       points[jj],
			}
		}(ii, ch)
	}

	centroidsMap := make(map[int][]Point)

	for ii := 0; ii < len(points); ii++ {
		result := <-ch
		centroidsMap[result.centroidInd] = append(centroidsMap[result.centroidInd], result.point)
	}

	return centroidsMap
}

func findClosest(p Point, centroids []Point) int {
	ch := make(chan Result)

	for ii := 0; ii < len(centroids); ii++ {
		go func(jj int, ch chan Result) {
			ch <- Result{
				centroidInd: jj,
				distance:    p.Distance(centroids[jj]),
			}
		}(ii, ch)
	}

	result := <-ch
	minDistance := result.distance
	minInd := result.centroidInd

	for ii := 1; ii < len(centroids); ii++ {
		result = <-ch
		if result.distance < minDistance {
			minDistance = result.distance
			minInd = result.centroidInd
		}
	}

	return minInd

}

func updateCentroids(oldCentroids []Point, centroidsMap map[int][]Point) []Point {

	ch := make(chan Result)
	for ii := 0; ii < len(oldCentroids); ii++ {
		go func(jj int, ch chan Result) {
			if len(centroidsMap[jj]) == 0 {
				ch <- Result{
					centroidInd: jj,
					point:       oldCentroids[jj],
				}
			} else {
				ch <- Result{
					centroidInd: jj,
					point:       averager(centroidsMap[jj]),
				}
			}
		}(ii, ch)
	}

	newCentroids := make([]Point, len(oldCentroids))
	for ii := 0; ii < len(oldCentroids); ii++ {
		result := <-ch
		newCentroids[result.centroidInd] = result.point
	}

	return newCentroids
}

func averager(points []Point) Point {
	newPoint := make(Point, len(points[0]))
	for _, point := range points {
		for jj, val := range point {
			newPoint[jj] += val
		}
	}
	for ii := range newPoint {
		newPoint[ii] /= float64(len(points))
	}
	return newPoint
}
