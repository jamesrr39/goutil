package dataprocessing

import (
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

	setA := Set{Points: []Point{Point{4.5, 6}, Point{6, 8}}}
	closestNeighbours := setClosestNeighboursWithSet(newPoint, nil, setA.Points, 0, numberOfNearestNeighbours)
	assert.Equal(t, []neighbour{neighbour{7.5, 0}, neighbour{10, 0}}, closestNeighbours)

	setB := Set{Points: []Point{Point{3, 4}}}
	closestNeighbours = setClosestNeighboursWithSet(newPoint, closestNeighbours, setB.Points, 1, numberOfNearestNeighbours)
	assert.Equal(t, []neighbour{neighbour{5, 1}, neighbour{7.5, 0}, neighbour{10, 0}}, closestNeighbours)
}

func Test_ClassifyPoint(t *testing.T) {
	const numberOfNearestNeighbours = 3

	setA := NewSet("set A", []Point{Point{6, 8}, Point{5, 8}, Point{5, 7.6}, Point{1, 5}, Point{10, 8}, Point{6, 32}, Point{6, 12}, Point{10, 8}})
	setB := NewSet("set B", []Point{Point{-6, -8}, Point{-3, -8}, Point{-4, -8}, Point{-6, -6}, Point{-6, -7}, Point{-6, -9}, Point{-10, -8}, Point{-16, -8}})

	classifier, err := NewKNNClassifier(numberOfNearestNeighbours, []Set{setA, setB})
	require.Nil(t, err)

	assert.Equal(t, "set A", classifier.ClassifyPoint(Point{10, 40}))
	assert.Equal(t, "set B", classifier.ClassifyPoint(Point{-10, -40}))
}
