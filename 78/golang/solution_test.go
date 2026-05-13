package problem78

import (
	"fmt"
	"sort"
	"testing"
)

// 78. Subsets
// Pattern: backtracking

func canonicalize(subs [][]int) []string {
	keys := make([]string, len(subs))
	for i, s := range subs {
		sorted := append([]int(nil), s...)
		sort.Ints(sorted)
		keys[i] = fmt.Sprintf("%v", sorted)
	}
	sort.Strings(keys)
	return keys
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"example 1", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"empty", []int{}, [][]int{{}}},
		{"single", []int{0}, [][]int{{}, {0}}},
		{"two", []int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canonicalize(subsets(tt.nums))
			want := canonicalize(tt.want)
			if !equal(got, want) {
				t.Errorf("subsets(%v) = %v; want %v", tt.nums, got, want)
			}
		})
	}
}
