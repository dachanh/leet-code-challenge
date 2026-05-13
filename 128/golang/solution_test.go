package problem128

import "testing"

// 128. Longest Consecutive Sequence
// Pattern: hashmap

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"example 2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"duplicates", []int{1, 2, 0, 1}, 3},
		{"negatives", []int{-3, -2, -1, 0, 1}, 5},
		{"disjoint", []int{1, 2, 100, 200, 300}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
