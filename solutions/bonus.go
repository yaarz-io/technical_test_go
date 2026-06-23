package exercises

// Unique returns a new slice with duplicate values removed, preserving the order
// of first occurrence.
// Example: [1, 2, 2, 3, 1] → [1, 2, 3]
func Unique(nums []int) []int {
	// seen tracks which values have already been encountered
	seen := make(map[int]bool)
	// result collects values in their original order, without duplicates
	result := []int{}
	// iterate over every number in the input slice
	for _, n := range nums {
		// only process the value if we haven't seen it before
		if !seen[n] {
			// mark this value as seen so future duplicates are skipped
			seen[n] = true
			// first occurrence: add it to the result
			result = append(result, n)
		}
	}
	// return the deduplicated slice
	return result
}

// Chunk splits nums into consecutive sub-slices of at most size elements.
// The last sub-slice may be shorter if len(nums) is not a multiple of size.
// Returns nil if nums is empty or size <= 0.
// Example: Chunk([]int{1, 2, 3, 4, 5}, 2) → [[1, 2], [3, 4], [5]]
func Chunk(nums []int, size int) [][]int {
	// nothing to split if the input is empty or the chunk size is invalid
	if len(nums) == 0 || size <= 0 {
		return nil
	}
	// result will hold each chunk as a sub-slice
	var result [][]int
	// advance the start index by size on every iteration
	for i := 0; i < len(nums); i += size {
		// end is the exclusive upper bound; min ensures we don't go past the slice length
		end := min(i+size, len(nums))
		// allocate a new slice so each chunk owns its own memory
		chunk := make([]int, end-i)
		// copy the elements from the source slice into the new chunk
		copy(chunk, nums[i:end])
		// append the finished chunk to the result
		result = append(result, chunk)
	}
	// return all chunks
	return result
}
