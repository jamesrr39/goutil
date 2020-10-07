package algorithms

type SearchResult int

const (
	// SearchResultFound indicates the desired result has been found and the search should terminate
	SearchResultFound SearchResult = iota
	// SearchResultGoLower indicates the search should look lower
	SearchResultGoLower
	// SearchResultGoHigher indicates the search should look higher
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
	if listSize == 0 {
		// handle special case
		return 0, false
	}

	// lowest and highest possible indexes
	lowerBound := 0
	upperBound := listSize - 1
	i := (listSize - 1) / 2 // start halfway through the set

	for iterations := 0; iterations < maxBinarySearchIterations; iterations++ {
		result := binarySearchFunc(i)
		switch result {
		case SearchResultFound:
			return i, true
		case SearchResultGoLower:
			// what we're searching for is below i
			// so set upperBound to be i - 1
			upperBound = i - 1
		case SearchResultGoHigher:
			// what we're searching for is above i
			// so set upperBound to be i + 1
			lowerBound = i + 1
		}
		if lowerBound > upperBound {
			// we've exhausted the search space without finding anything
			return i, false
		}
		i = (lowerBound + upperBound) / 2

	}
	panic("maxBinarySearchIterations reached")
}
