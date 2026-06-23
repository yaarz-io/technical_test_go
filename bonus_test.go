package exercises

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{5, 5, 5}, []int{5}},
	}
	for _, tt := range tests {
		got := Unique(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Unique(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		input []int
		size  int
		want  [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{}, 2, nil},
		{[]int{1, 2, 3}, 0, nil},
	}
	for _, tt := range tests {
		got := Chunk(tt.input, tt.size)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Chunk(%v, %d) = %v, want %v", tt.input, tt.size, got, tt.want)
		}
	}
}
