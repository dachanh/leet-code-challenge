package problem347

import (
	"sort"
	"testing"
)

// 347. Top K Frequent Elements
// Pattern: hashmap + bucket-sort

func sortedEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	ac := append([]int(nil), a...)
	bc := append([]int(nil), b...)
	sort.Ints(ac)
	sort.Ints(bc)
	for i := range ac {
		if ac[i] != bc[i] {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example 1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"example 2", []int{1}, 1, []int{1}},
		{"all same freq", []int{1, 2, 3, 4}, 2, []int{1, 2}},
		{"k = unique", []int{1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
		{"with negatives", []int{-1, -1, -2, -3, -3, -3}, 2, []int{-3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !sortedEqual(got, tt.want) {
				t.Errorf("topKFrequent(%v, %d) = %v; want set %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
