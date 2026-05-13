package problem215

import "testing"

// 215. Kth Largest Element in an Array
// Pattern: heap

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"example 1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"example 2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"single", []int{1}, 1, 1},
		{"all same", []int{2, 2, 2, 2}, 3, 2},
		{"k=1 largest", []int{1, 2, 3, 4, 5}, 1, 5},
		{"k=n smallest", []int{1, 2, 3, 4, 5}, 5, 1},
		{"negatives", []int{-1, -2, -3, -4}, 2, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.nums, tt.k); got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
