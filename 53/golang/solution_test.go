package problem53

import "testing"

// 53. Maximum Subarray
// Pattern: kadane

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"example 2", []int{1}, 1},
		{"example 3", []int{5, 4, -1, 7, 8}, 23},
		{"all negatives", []int{-3, -1, -2}, -1},
		{"single neg", []int{-5}, -5},
		{"all positive", []int{1, 2, 3, 4}, 10},
		{"zeros", []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
