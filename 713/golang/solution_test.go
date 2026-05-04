package problem713

import "testing"

// 713. Subarray Product Less Than K
// Pattern: sliding-window
// Signature: func numSubarrayProductLessThanK(nums []int, k int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1: classic LeetCode case",
			nums: []int{10, 5, 2, 6},
			k:    100,
			want: 8,
		},
		{
			name: "example 2: k=0 yields no subarrays",
			nums: []int{1, 2, 3},
			k:    0,
			want: 0,
		},
		{
			name: "k=1 with positive ints yields none",
			nums: []int{1, 2, 3},
			k:    1,
			want: 0,
		},
		{
			name: "single element less than k",
			nums: []int{2},
			k:    5,
			want: 1,
		},
		{
			name: "single element equal to k (excluded since strict)",
			nums: []int{5},
			k:    5,
			want: 0,
		},
		{
			name: "single element greater than k",
			nums: []int{10},
			k:    5,
			want: 0,
		},
		{
			name: "all ones with large k",
			nums: []int{1, 1, 1, 1},
			k:    2,
			want: 10,
		},
		{
			name: "very large k counts every subarray",
			nums: []int{1, 2, 3, 4},
			k:    1000000,
			want: 10,
		},
		{
			name: "mix where some elements alone exceed k",
			nums: []int{100, 1, 1, 1},
			k:    50,
			want: 6,
		},
		{
			name: "trailing element exceeds k",
			nums: []int{1, 2, 3, 100},
			k:    50,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSubarrayProductLessThanK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("numSubarrayProductLessThanK(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
