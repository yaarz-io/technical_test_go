package exercises

import (
	"errors"
	"strings"
)

// Product represents an item in a store catalog.
type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	InStock  bool
}

// Sum returns the sum of all integers in nums.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Filter returns a new slice containing only the elements for which keep returns true.
// The order of elements is preserved.
func Filter(nums []int, keep func(int) bool) []int {
	result := []int{}
	for _, n := range nums {
		if keep(n) {
			result = append(result, n)
		}
	}
	return result
}

// CountVowels returns the number of vowels (a, e, i, o, u) in s (case-insensitive).
func CountVowels(s string) int {
	count := 0
	for _, r := range strings.ToLower(s) {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

// SafeDivide returns a / b.
// It returns an error if b is zero.
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// GroupByCategory returns a map where each key is a category name and the value
// is the list of products belonging to that category.
func GroupByCategory(products []Product) map[string][]Product {
	result := make(map[string][]Product)
	for _, p := range products {
		result[p.Category] = append(result[p.Category], p)
	}
	return result
}
