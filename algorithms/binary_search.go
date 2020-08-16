package algorithms

type SearchResult int

const (
	SearchResultFound SearchResult = iota
	SearchResultGoLower
	SearchResultGoHigher
)

// BinarySearchFunc is a caller-supplied function that tells the Binary Search function to go higher or lower
type BinarySearchFunc func(i int) SearchResult

const (
	// maxBinarySearchIterations is the maximum amount of iterations allowed before a panic is created.
	// since the worst case time complexity of binary search is log(2)n, n = 100000 allows for a huge amount of items to be tested
	maxBinarySearchIterations = 100000
)

// BinarySearch performs a binary search on a given list size and binary search function
// It returns the index of the last value tested it could find, and a boolean indicating whether the value was found exactly
func BinarySearch(listSize int, binarySearchFunc BinarySearchFunc) (int, bool) {
	lowerBound := 0
	upperBound := listSize - 1
	i := (listSize - 1) / 2 // start halfway through the set

	for iterations := 0; iterations < maxBinarySearchIterations; iterations++ {
		result := binarySearchFunc(i)
		switch result {
		case SearchResultFound:
			return i, true
		case SearchResultGoLower:
			upperBound = i
		case SearchResultGoHigher:
			lowerBound = i
		}
		i = (lowerBound + upperBound) / 2
		if i == lowerBound || i == upperBound {
			return i, false
		}
	}
	panic("maxBinarySearchIterations reached")
}
