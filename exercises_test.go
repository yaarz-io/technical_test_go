package exercises

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 1}, 0},
		{[]int{10, 20, 30}, 60},
	}
	for _, tt := range tests {
		got := Sum(tt.input)
		if got != tt.want {
			t.Errorf("Sum(%v) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }

	tests := []struct {
		input []int
		fn    func(int) bool
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, isEven, []int{2, 4}},
		{[]int{2, 4, 6}, isEven, []int{2, 4, 6}},
		{[]int{1, 3, 5}, isEven, []int{}},
		{[]int{}, isEven, []int{}},
		{[]int{-1, 0, 1, 2}, isPositive, []int{1, 2}},
	}
	for _, tt := range tests {
		got := Filter(tt.input, tt.fn)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Filter(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"hello", 2},
		{"AEIOU", 5},
		{"rhythm", 0},
		{"Hello World", 3},
	}
	for _, tt := range tests {
		got := CountVowels(tt.input)
		if got != tt.want {
			t.Errorf("CountVowels(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		a, b    float64
		want    float64
		wantErr bool
	}{
		{10, 2, 5, false},
		{9, 3, 3, false},
		{0, 5, 0, false},
		{10, 0, 0, true},
		{-6, 2, -3, false},
	}
	for _, tt := range tests {
		got, err := SafeDivide(tt.a, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("SafeDivide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("SafeDivide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestGroupByCategory(t *testing.T) {
	products := []Product{
		{ID: 1, Name: "Laptop", Category: "Electronics", Price: 999.99, InStock: true},
		{ID: 2, Name: "Phone", Category: "Electronics", Price: 599.99, InStock: false},
		{ID: 3, Name: "Desk", Category: "Furniture", Price: 249.99, InStock: true},
		{ID: 4, Name: "Chair", Category: "Furniture", Price: 149.99, InStock: true},
		{ID: 5, Name: "Monitor", Category: "Electronics", Price: 399.99, InStock: true},
	}

	got := GroupByCategory(products)

	if len(got) != 2 {
		t.Errorf("GroupByCategory() returned %d categories, want 2", len(got))
	}
	if len(got["Electronics"]) != 3 {
		t.Errorf("Electronics has %d products, want 3", len(got["Electronics"]))
	}
	if len(got["Furniture"]) != 2 {
		t.Errorf("Furniture has %d products, want 2", len(got["Furniture"]))
	}
}
