package binaryx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipEndiannessInPlace(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name     string
		args     args
		expected []byte
	}{
		{
			name: "odd number of bytes",
			args: args{
				s: []byte{'a', 'b', 'c', 'd', 'e'},
			},
			expected: []byte{'e', 'd', 'c', 'b', 'a'},
		},
		{
			name: "even number of bytes",
			args: args{
				s: []byte{'a', 'b', 'c', 'd', 'e', 'f'},
			},
			expected: []byte{'f', 'e', 'd', 'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlipEndiannessInPlace(tt.args.s)
			assert.Equal(t, tt.expected, tt.args.s)
		})
	}
}

func TestFlipEndiannessInNewSlice(t *testing.T) {
	type args struct {
		old []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				old: []byte{'a', 'b', 'c', 'd', 'e'},
			},
			want: []byte{'e', 'd', 'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipEndiannessInNewSlice(tt.args.old); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlipEndiannessInNewSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
