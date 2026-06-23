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
	// start the accumulator at zero
	total := 0
	// iterate over every number in the slice
	for _, n := range nums {
		// add the current number to the running total
		total += n
	}
	// return the final sum
	return total
}

// Filter returns a new slice containing only the elements for which keep returns true.
// The order of elements is preserved.
func Filter(nums []int, keep func(int) bool) []int {
	// initialise an empty slice to collect results (not nil, so it serialises cleanly)
	result := []int{}
	// iterate over every number in the input slice
	for _, n := range nums {
		// call the predicate: only keep the element if it returns true
		if keep(n) {
			// append the accepted element to the result slice
			result = append(result, n)
		}
	}
	// return the filtered slice
	return result
}

// CountVowels returns the number of vowels (a, e, i, o, u) in s (case-insensitive).
func CountVowels(s string) int {
	// start the counter at zero
	count := 0
	// convert to lowercase once so the switch below only needs to handle lowercase letters
	for _, r := range strings.ToLower(s) {
		// check whether the current rune is one of the five vowels
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			// it is a vowel: increment the counter
			count++
		}
	}
	// return the total number of vowels found
	return count
}

// SafeDivide returns a / b.
// It returns an error if b is zero.
func SafeDivide(a, b float64) (float64, error) {
	// guard against division by zero, which is undefined
	if b == 0 {
		// return a descriptive error; the zero float is the conventional empty value
		return 0, errors.New("division by zero")
	}
	// b is non-zero: perform the division and signal no error
	return a / b, nil
}

// GroupByCategory returns a map where each key is a category name and the value
// is the list of products belonging to that category.
func GroupByCategory(products []Product) map[string][]Product {
	// create an empty map; keys are category strings, values are slices of products
	result := make(map[string][]Product)
	// iterate over every product in the input slice
	for _, p := range products {
		// append the product to its category's slice (the slice is created automatically on first use)
		result[p.Category] = append(result[p.Category], p)
	}
	// return the populated map
	return result
}
