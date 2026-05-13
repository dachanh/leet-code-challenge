package problem560

import "testing"

// 560. Subarray Sum Equals K
// Pattern: prefix-sum

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example 1", []int{1, 1, 1}, 2, 2},
		{"example 2", []int{1, 2, 3}, 3, 2},
		{"single match", []int{3}, 3, 1},
		{"single no match", []int{3}, 4, 0},
		{"with negatives", []int{1, -1, 0}, 0, 3},
		{"zeros", []int{0, 0, 0}, 0, 6},
		{"target zero", []int{1, -1, 1, -1}, 0, 4},
		{"all negatives", []int{-1, -1, -2}, -3, 1},
		{"no subarray", []int{1, 2, 3}, 7, 0},
		{"long run", []int{1, 1, 1, 1, 1}, 2, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("subarraySum(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
