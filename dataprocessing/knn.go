package dataprocessing

import (
	"errors"
	"math"
)

var (
	ErrEvenNumberOfNearestNeighbours = errors.New("gave an even number of nearest neighbours, but it should be odd, to avoid a tie")
)

// https://www.analyticsvidhya.com/blog/2018/03/introduction-k-neighbours-algorithm-clustering/
// http://www.statsoft.com/Textbook/k-Nearest-Neighbors#predictions

type Point struct {
	X float64
	Y float64
}

type KNNClassifier struct {
	NumberOfNearestNeighbours int
	Sets                      []Set
}

func NewKNNClassifier(numberOfNearestNeighbours int, sets []Set) (*KNNClassifier, error) {
	if numberOfNearestNeighbours%2 == 0 {
		return nil, ErrEvenNumberOfNearestNeighbours
	}
	return &KNNClassifier{numberOfNearestNeighbours, sets}, nil
}

type neighbour struct {
	Distance float64
	setIndex int
}

type Set struct {
	Name   string
	Points []Point
}

func NewSet(name string, points []Point) Set {
	return Set{name, points}
}

func (c *KNNClassifier) ClassifyPoint(point Point) string {
	var closestNeighbours []neighbour
	for i := 0; i < len(c.Sets); i++ {
		closestNeighbours = setClosestNeighboursWithSet(point, closestNeighbours, c.Sets[i].Points, i, c.NumberOfNearestNeighbours)
	}

	nearbySetCounts := make(map[int]int)
	for _, closeNeighbour := range closestNeighbours {
		nearbySetCounts[closeNeighbour.setIndex]++
	}

	closestSetID := -1
	closestSetCount := 0
	for setID, count := range nearbySetCounts {
		if count > closestSetCount {
			closestSetID = setID
		}
	}

	return c.Sets[closestSetID].Name
}

func setClosestNeighboursWithSet(point Point, existingClosestNeighbours []neighbour, set []Point, setIndex, numberOfNearestNeighbours int) []neighbour {
	for _, setPoint := range set {
		distance := distanceFromPoint(point, setPoint)

		// replace an existing item in the set
		for i := 0; i < numberOfNearestNeighbours; i++ {
			if i < (len(existingClosestNeighbours)) && distance > existingClosestNeighbours[i].Distance {
				continue
			}

			itemCountAfter := numberOfNearestNeighbours - 1
			if itemCountAfter > (len(existingClosestNeighbours) - 1) {
				itemCountAfter = len(existingClosestNeighbours)
			}

			itemsBefore := existingClosestNeighbours[:i]
			itemsAfter := existingClosestNeighbours[i:itemCountAfter]

			existingClosestNeighbours = append(
				itemsBefore,
				append(
					[]neighbour{neighbour{distance, setIndex}},
					itemsAfter...,
				)...,
			)
			break
		}
	}

	return existingClosestNeighbours
}

func distanceFromPoint(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow(point1.X-point2.X, 2) + math.Pow(point1.Y-point2.Y, 2))
}
