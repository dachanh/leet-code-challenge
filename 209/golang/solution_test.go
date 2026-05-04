package problem209

import "testing"

// 209. Minimum Size Subarray Sum
// Pattern: sliding-window
// Signature: func minSubArrayLen(target int, nums []int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "example 1: classic LeetCode case",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "example 2: single element meets target",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "example 3: total sum less than target",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "empty input",
			target: 5,
			nums:   []int{},
			want:   0,
		},
		{
			name:   "single element less than target",
			target: 100,
			nums:   []int{50},
			want:   0,
		},
		{
			name:   "single element equals target",
			target: 5,
			nums:   []int{5},
			want:   1,
		},
		{
			name:   "entire array needed",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			want:   5,
		},
		{
			name:   "target zero (any prefix works)",
			target: 0,
			nums:   []int{1, 2, 3},
			want:   1,
		},
		{
			name:   "large numbers, small window",
			target: 100,
			nums:   []int{50, 60, 70, 80},
			want:   2,
		},
		{
			name:   "needs a long run of 1s",
			target: 6,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d; want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
