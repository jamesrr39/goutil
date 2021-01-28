package algorithms

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/jamesrr39/goutil/binaryx"
	"github.com/jamesrr39/goutil/errorsx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	listEvenLength := []int{1, 4, 6, 10, 34, 80}
	listOddLength := []int{1, 4, 6, 10, 34, 80, 102}
	type args struct {
		listSize         int
		binarySearchFunc BinarySearchFunc
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 SearchResult
	}{
		{
			name: "example found list even length",
			args: args{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := 6
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  2,
			want1: SearchResultFound,
		}, {
			name: "example not found list even length",
			args: args{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := 7
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  3,
			want1: SearchResultGoLower,
		}, {
			name: "example found list odd length",
			args: args{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := 6
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  2,
			want1: SearchResultFound,
		}, {
			name: "example not found list even length",
			args: args{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := 7
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  2,
			want1: SearchResultGoHigher,
		}, {
			name: "example not found, value too low, list even length",
			args: args{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := -1
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  0,
			want1: SearchResultGoLower,
		}, {
			name: "example not found, value too high, list even length",
			args: args{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) SearchResult {
					wantedValue := 700000
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound
					}

					if val > wantedValue {
						return SearchResultGoLower
					}

					return SearchResultGoHigher
				},
			},
			want:  5,
			want1: SearchResultGoHigher,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearch(tt.args.listSize, tt.args.binarySearchFunc)
			if got != tt.want {
				t.Errorf("BinarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_17(t *testing.T) {
	dataSet := []uint64{3, 5, 6, 7, 12, 14, 18, 22, 25, 30, 34, 42, 51, 63, 64, 65, 101}
	var items [][]byte
	for _, s := range dataSet {
		b := binaryx.LittleEndianPutUint64(s)
		items = append(items, b)
	}

	// setup search function
	var binarySearchErr error
	makeFunc := func(key []byte) BinarySearchFunc {
		return func(i int) SearchResult {
			thisKey := items[i]

			if bytes.Equal(thisKey, key) {
				return SearchResultFound
			}

			thisIsLarger, err := isKey1GreaterThanKey2CompareInt64Func(thisKey, key)
			if err != nil {
				binarySearchErr = err
				return SearchResultFound
			}

			if thisIsLarger {
				// this one is greater than the searched for. So look lower
				return SearchResultGoLower
			}

			return SearchResultGoHigher
		}
	}

	for i, item := range items {
		t.Run(fmt.Sprintf("find %d, idx %d", dataSet[i], i), func(t *testing.T) {
			idx, lastResult := BinarySearch(len(dataSet), makeFunc(item))
			require.NoError(t, binarySearchErr)
			assert.Equal(t, SearchResultFound, lastResult)
			assert.Equal(t, i, idx)
		})
	}
}

func isKey1GreaterThanKey2CompareInt64Func(key1, key2 []byte) (bool, errorsx.Error) {
	val1 := binary.LittleEndian.Uint64(key1)
	val2 := binary.LittleEndian.Uint64(key2)
	return val1 > val2, nil
}
