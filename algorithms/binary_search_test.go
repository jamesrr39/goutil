package algorithms

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	listEvenLength := []int{1, 4, 6, 10, 34, 80}
	listOddLength := []int{1, 4, 6, 10, 34, 80, 102}
	type args[T numberType] struct {
		listSize         T
		binarySearchFunc BinarySearchFunc[T]
	}
	tests := []struct {
		name  string
		args  args[int]
		want  BinarySearchResult[int]
		want1 error
	}{
		{
			name: "example found list even length",
			args: args[int]{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := 6
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 2,
				LastResult:       SearchResultFound,
			},
			want1: nil,
		}, {
			name: "example not found list even length",
			args: args[int]{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := 7
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 3,
				LastResult:       SearchResultGoLower,
			},
			want1: nil,
		}, {
			name: "example found list odd length",
			args: args[int]{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := 6
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 2,
				LastResult:       SearchResultFound,
			},
			want1: nil,
		}, {
			name: "example not found list even length",
			args: args[int]{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := 7
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 2,
				LastResult:       SearchResultGoHigher,
			},
			want1: nil,
		}, {
			name: "example not found, value too low, list even length",
			args: args[int]{
				listSize: len(listOddLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := -1
					val := listOddLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 0,
				LastResult:       SearchResultGoLower,
			},
			want1: nil,
		}, {
			name: "example not found, value too high, list even length",
			args: args[int]{
				listSize: len(listEvenLength),
				binarySearchFunc: func(i int) (SearchResult, error) {
					wantedValue := 700000
					val := listEvenLength[i]
					if val == wantedValue {
						return SearchResultFound, nil
					}

					if val > wantedValue {
						return SearchResultGoLower, nil
					}

					return SearchResultGoHigher, nil
				},
			},
			want: BinarySearchResult[int]{
				LastIndexChecked: 5,
				LastResult:       SearchResultGoHigher,
			},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearch(tt.args.listSize, tt.args.binarySearchFunc, nil)
			if got != tt.want {
				t.Errorf("BinarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
