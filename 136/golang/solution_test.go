package problem136

import "testing"

// 136. Single Number
// Pattern: bit-manipulation

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{2, 2, 1}, 1},
		{"example 2", []int{4, 1, 2, 1, 2}, 4},
		{"example 3", []int{1}, 1},
		{"negatives", []int{-1, -1, 5}, 5},
		{"zero solo", []int{0, 1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
