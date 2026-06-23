package exercises

import (
	"strings"
	"unicode"
)

// Unique returns a new slice with duplicate values removed, preserving the order
// of first occurrence.
// Example: [1, 2, 2, 3, 1] → [1, 2, 3]
func Unique(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}

// ReverseWords returns s with the order of words reversed.
// Words are separated by whitespace; leading/trailing spaces are ignored.
// Example: "Hello World" → "World Hello"
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// IsPalindrome reports whether s reads the same forwards and backwards.
// Only alphanumeric characters are considered; comparison is case-insensitive.
// Example: "A man a plan a canal Panama" → true
func IsPalindrome(s string) bool {
	var chars []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			chars = append(chars, r)
		}
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}

// MostExpensiveByCategory returns a map where each key is a category and the value
// is the most expensive product in that category (by Price).
// If two products share the same price, either may be returned.
func MostExpensiveByCategory(products []Product) map[string]Product {
	result := make(map[string]Product)
	for _, p := range products {
		current, exists := result[p.Category]
		if !exists || p.Price > current.Price {
			result[p.Category] = p
		}
	}
	return result
}

// Chunk splits nums into consecutive sub-slices of at most size elements.
// The last sub-slice may be shorter if len(nums) is not a multiple of size.
// Returns nil if nums is empty or size <= 0.
// Example: Chunk([]int{1, 2, 3, 4, 5}, 2) → [[1, 2], [3, 4], [5]]
func Chunk(nums []int, size int) [][]int {
	if len(nums) == 0 || size <= 0 {
		return nil
	}
	var result [][]int
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > len(nums) {
			end = len(nums)
		}
		chunk := make([]int, end-i)
		copy(chunk, nums[i:end])
		result = append(result, chunk)
	}
	return result
}
