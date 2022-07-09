package dataprocessing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewKNNClassifier(t *testing.T) {
	_, err := NewKNNClassifier(2, nil)
	assert.Equal(t, ErrEvenNumberOfNearestNeighbours, err)
}

func Test_distanceFromPoint(t *testing.T) {
	assert.Equal(t, float64(5), distanceFromPoint(Point{X: 4, Y: 5}, Point{X: 1, Y: 1}))

	assert.Equal(t, float64(100), distanceFromPoint(Point{X: -50, Y: -40}, Point{X: 30, Y: 20}))
}

func Test_setClosestNeighboursWithSet(t *testing.T) {
	const numberOfNearestNeighbours = 3
	newPoint := Point{0, 0}

	setA := Set{Points: []Point{{4.5, 6}, {6, 8}}}
	closestNeighbours := setClosestNeighboursWithSet(newPoint, nil, setA.Points, 0, numberOfNearestNeighbours)
	assert.Equal(t, []neighbour{{7.5, 0}, {10, 0}}, closestNeighbours)

	setB := Set{Points: []Point{{3, 4}}}
	closestNeighbours = setClosestNeighboursWithSet(newPoint, closestNeighbours, setB.Points, 1, numberOfNearestNeighbours)
	assert.Equal(t, []neighbour{{5, 1}, {7.5, 0}, {10, 0}}, closestNeighbours)
}

func Test_ClassifyPoint(t *testing.T) {
	const numberOfNearestNeighbours = 3

	setA := NewSet("set A", []Point{{6, 8}, {5, 8}, {5, 7.6}, {1, 5}, {10, 8}, {6, 32}, {6, 12}, {10, 8}})
	setB := NewSet("set B", []Point{{-6, -8}, {-3, -8}, {-4, -8}, {-6, -6}, {-6, -7}, {-6, -9}, {-10, -8}, {-16, -8}})

	classifier, err := NewKNNClassifier(numberOfNearestNeighbours, []Set{setA, setB})
	require.Nil(t, err)

	assert.Equal(t, "set A", classifier.ClassifyPoint(Point{10, 40}))
	assert.Equal(t, "set B", classifier.ClassifyPoint(Point{-10, -40}))
}

func ExampleNewKNNClassifier() {
	// X = weight kg, Y = number of passengers
	bicycleSet := NewSet("bicycle", []Point{{X: 10, Y: 1}, {X: 15, Y: 1}, {X: 10, Y: 1}, {X: 11, Y: 1}, {X: 18, Y: 1}, {X: 20, Y: 1}})
	carSet := NewSet("car", []Point{{X: 1000, Y: 4}, {X: 1500, Y: 5}, {X: 900, Y: 2}, {X: 1300, Y: 5}, {X: 2000, Y: 4}})

	const numberOfNearestNeighbours = 3

	classifier, err := NewKNNClassifier(numberOfNearestNeighbours, []Set{bicycleSet, carSet})
	if err != nil {
		panic(err)
	}

	newBike := Point{
		X: 25,
		Y: 2, // with child seat
	}
	newCar := Point{
		X: 1300,
		Y: 4,
	}

	newBikeLabel := classifier.ClassifyPoint(newBike)
	newCarLabel := classifier.ClassifyPoint(newCar)

	fmt.Printf("'newBike' classified as: %s\n", newBikeLabel)
	fmt.Printf("'newCar' classified as: %s\n", newCarLabel)

	// Output:
	// 'newBike' classified as: bicycle
	// 'newCar' classified as: car
}
