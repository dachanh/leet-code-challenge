package problem33

import "testing"

// 33. Search in Rotated Sorted Array
// Pattern: binary-search

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"example 1", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"example 2", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"example 3", []int{1}, 0, -1},
		{"not rotated", []int{1, 2, 3, 4, 5}, 4, 3},
		{"rotated once", []int{5, 1, 2, 3, 4}, 1, 1},
		{"target at pivot", []int{6, 7, 1, 2, 3, 4, 5}, 1, 2},
		{"target at start", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"target at end", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"two element rotated", []int{3, 1}, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
