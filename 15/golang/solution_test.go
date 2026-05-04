package problem15

import (
	"sort"
	"testing"
)

// 15. 3Sum
// Pattern: fix+scan
// Signature: func threeSum(nums []int) [][]int
//
// Note: LeetCode accepts triplets in any order. We normalize by sorting
// each triplet, then sorting the outer slice, before comparison.

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1: classic mix of positives/negatives",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2: no triplet sums to zero",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3: all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "two elements only",
			nums: []int{1, -1},
			want: [][]int{},
		},
		{
			name: "all positives, no solution",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "all negatives, no solution",
			nums: []int{-1, -2, -3, -4},
			want: [][]int{},
		},
		{
			name: "many duplicates collapse to one triplet",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "duplicates yielding multiple triplets",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "larger mixed input",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !equalTriplets(got, tt.want) {
				t.Errorf("threeSum(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func equalTriplets(a, b [][]int) bool {
	na := normalizeTriplets(a)
	nb := normalizeTriplets(b)
	if len(na) != len(nb) {
		return false
	}
	for i := range na {
		if len(na[i]) != len(nb[i]) {
			return false
		}
		for j := range na[i] {
			if na[i][j] != nb[i][j] {
				return false
			}
		}
	}
	return true
}

func normalizeTriplets(in [][]int) [][]int {
	out := make([][]int, len(in))
	for i, t := range in {
		cp := append([]int(nil), t...)
		sort.Ints(cp)
		out[i] = cp
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return len(out[i]) < len(out[j])
	})
	return out
}
