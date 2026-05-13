package problem16

import "testing"

// 16. 3Sum Closest
// Pattern: fix+scan (two pointers)
// Signature: func threeSumClosest(nums []int, target int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1: classic case",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "example 2: all zeros",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "exact match possible",
			nums:   []int{1, 1, 1, 0},
			target: -100,
			want:   2,
		},
		{
			name:   "negative target",
			nums:   []int{-1, -1, -1, -1},
			target: -3,
			want:   -3,
		},
		{
			name:   "minimum length triplet",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   6,
		},
		{
			name:   "minimum length, target far away",
			nums:   []int{-5, -2, 0},
			target: 100,
			want:   -7,
		},
		{
			name:   "duplicates with positive target",
			nums:   []int{1, 1, 1, 1, 1, 1},
			target: 4,
			want:   3,
		},
		{
			name:   "mixed signs, target zero",
			nums:   []int{-3, -2, -5, 3, -4},
			target: -1,
			want:   -2,
		},
		{
			name:   "large positive numbers",
			nums:   []int{10, 20, 30, 40, 50},
			target: 100,
			want:   100,
		},
		{
			name:   "tie: two sums equally close — either is acceptable",
			nums:   []int{0, 1, 2},
			target: 3,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumClosest(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
