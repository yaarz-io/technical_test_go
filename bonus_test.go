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

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello World", "World Hello"},
		{"one two three", "three two one"},
		{"single", "single"},
		{"", ""},
		{"  extra  spaces  ", "spaces extra"},
	}
	for _, tt := range tests {
		got := ReverseWords(tt.input)
		if got != tt.want {
			t.Errorf("ReverseWords(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man a plan a canal Panama", true},
		{"", true},
		{"Was it a car or a cat I saw", true},
		{"Not a palindrome", false},
	}
	for _, tt := range tests {
		got := IsPalindrome(tt.input)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestMostExpensiveByCategory(t *testing.T) {
	products := []Product{
		{ID: 1, Name: "Laptop", Category: "Electronics", Price: 999.99, InStock: true},
		{ID: 2, Name: "Phone", Category: "Electronics", Price: 599.99, InStock: false},
		{ID: 3, Name: "Desk", Category: "Furniture", Price: 249.99, InStock: true},
		{ID: 4, Name: "Chair", Category: "Furniture", Price: 149.99, InStock: true},
		{ID: 5, Name: "Monitor", Category: "Electronics", Price: 399.99, InStock: true},
	}

	got := MostExpensiveByCategory(products)

	if got["Electronics"].ID != 1 {
		t.Errorf("most expensive Electronics: got %s ($%.2f), want Laptop ($999.99)",
			got["Electronics"].Name, got["Electronics"].Price)
	}
	if got["Furniture"].ID != 3 {
		t.Errorf("most expensive Furniture: got %s ($%.2f), want Desk ($249.99)",
			got["Furniture"].Name, got["Furniture"].Price)
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
