package algorithms

import (
	"bytes"
	"errors"
	"fmt"
)

//go:generate stringer -type=SearchResult
type SearchResult int

const (
	SearchResultUnknown SearchResult = iota
	// SearchResultFound indicates the desired result has been found and the search should terminate
	SearchResultFound
	// SearchResultGoLower indicates the search should look lower
	SearchResultGoLower
	// SearchResultGoHigher indicates the search should look higher
	SearchResultGoHigher
)

var (
	// This can mean that a list with 0 items was passed in.
	ErrSearchResultInvalid = errors.New("ErrSearchResultInvalid")
)

// BinarySearchFunc is a caller-supplied function that tells the Binary Search function to go higher or lower
type BinarySearchFunc[T numberType] func(i T) (SearchResult, error)

func makeDefaultBinarySearchOpts() BinarySearchOptions {
	return BinarySearchOptions{
		MaxIterations: 100000,
	}
}

type BinarySearchOptions struct {
	MaxIterations int64
}

func mergeOpts(passedInOptions *BinarySearchOptions) BinarySearchOptions {
	if passedInOptions == nil {
		return makeDefaultBinarySearchOpts()
	}

	opts := makeDefaultBinarySearchOpts()

	if passedInOptions.MaxIterations != 0 {
		opts.MaxIterations = passedInOptions.MaxIterations
	}

	return opts
}

type BinarySearchResult[T numberType] struct {
	LastIndexChecked T
	LastResult       SearchResult
}

type numberType interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

// BinarySearch performs a binary search on a given list size and binary search function
// please note, the exact value you search for may not be found - please check the BinarySearchResult value
func BinarySearch[T numberType](listSize T, binarySearchFunc BinarySearchFunc[T], passedInOptions *BinarySearchOptions) (BinarySearchResult[T], error) {
	options := mergeOpts(passedInOptions)

	if int(listSize) == 0 {
		// handle special case
		return BinarySearchResult[T]{}, ErrSearchResultInvalid
	}

	// lowest and highest possible indexes
	var lowerBound T = 0
	upperBound := listSize - 1
	i := (listSize - 1) / 2 // start halfway through the set

	for iterations := 0; iterations < int(options.MaxIterations); iterations++ {
		result, err := binarySearchFunc(i)
		if err != nil {
			return BinarySearchResult[T]{}, err
		}

		switch result {
		case SearchResultFound:
			return BinarySearchResult[T]{
				LastIndexChecked: i,
				LastResult:       result,
			}, nil
			// return i, result
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
			return BinarySearchResult[T]{
				LastIndexChecked: i,
				LastResult:       result,
			}, nil
		}
		i = (lowerBound + upperBound) / 2

	}
	panic(fmt.Sprintf("maxBinarySearchIterations (%d) reached", options.MaxIterations))
}

// CreateByteComparatorFunc creates a comparator function for keys of type []byte
func CreateByteComparatorFunc[T numberType](needle []byte, getListValueAtIndex func(index T) []byte) BinarySearchFunc[T] {
	return func(index T) (SearchResult, error) {
		compareResult := bytes.Compare(needle, getListValueAtIndex(index))
		switch compareResult {
		case 0:
			return SearchResultFound, nil
		case 1:
			return SearchResultGoHigher, nil
		case -1:
			return SearchResultGoLower, nil
		}
		return 0, fmt.Errorf("unexpected bytes.Compare result: %d", compareResult)
	}
}
