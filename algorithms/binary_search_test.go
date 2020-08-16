package algorithms

import "testing"

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
		want1 bool
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
			want1: true,
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
			want:  2,
			want1: false,
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
			want1: true,
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
			want1: false,
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
			want1: false,
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
			want:  4,
			want1: false,
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
