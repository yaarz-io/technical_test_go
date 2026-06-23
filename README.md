# Go Technical Test

Implement the functions in `exercises.go` and (optionally) `bonus.go`. Tests are already written — your goal is to make them pass.

**Estimated time:** 20–30 minutes for the core exercises.

## Getting started

```bash
# Run all tests
go test ./...

# Run a single test by name (useful while implementing one function at a time)
go test -run TestSum .
go test -run TestFilter .

# Run with details on which tests pass/fail
go test -v .
```

Work function by function: implement one, run its test with `-run`, then move to the next. You are done when `go test ./...` shows `ok`.

## Exercises (`exercises.go`)

Implement these 5 functions:

| Function | Description |
|---|---|
| `Sum(nums []int) int` | Return the sum of all elements |
| `Filter(nums []int, keep func(int) bool) []int` | Return only elements for which `keep` is true |
| `CountVowels(s string) int` | Count vowels (a, e, i, o, u), case-insensitive |
| `SafeDivide(a, b float64) (float64, error)` | Divide `a` by `b`; return an error if `b` is zero |
| `GroupByCategory(products []Product) map[string][]Product` | Group a slice of products by their category |

## Bonus (`bonus.go`)

If you finish early, implement these:

| Function | Description |
|---|---|
| `Unique(nums []int) []int` | Remove duplicates, preserving order of first occurrence |
| `Chunk(nums []int, size int) [][]int` | Split a slice into consecutive sub-slices of `size` elements |

## Rules

- Do not modify `*_test.go` files.
- The `solutions/` directory is off-limits.
- You may use the Go standard library freely.
