package exercises

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
	// TODO: implement
	panic("not implemented")
}

// Filter returns a new slice containing only the elements for which keep returns true.
// The order of elements is preserved.
func Filter(nums []int, keep func(int) bool) []int {
	// TODO: implement
	panic("not implemented")
}

// CountVowels returns the number of vowels (a, e, i, o, u) in s (case-insensitive).
func CountVowels(s string) int {
	// TODO: implement
	panic("not implemented")
}

// SafeDivide returns a / b.
// It returns an error if b is zero.
func SafeDivide(a, b float64) (float64, error) {
	// TODO: implement
	panic("not implemented")
}

// GroupByCategory returns a map where each key is a category name and the value
// is the list of products belonging to that category.
func GroupByCategory(products []Product) map[string][]Product {
	// TODO: implement
	panic("not implemented")
}
