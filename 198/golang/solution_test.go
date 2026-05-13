package problem198

import "testing"

// 198. House Robber
// Pattern: dp

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{1, 2, 3, 1}, 4},
		{"example 2", []int{2, 7, 9, 3, 1}, 12},
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"two", []int{2, 1}, 2},
		{"alternating big", []int{10, 1, 10, 1, 10}, 30},
		{"all zero", []int{0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
