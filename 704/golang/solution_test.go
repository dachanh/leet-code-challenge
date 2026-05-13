package problem704

import "testing"

// 704. Binary Search
// Pattern: binary-search

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"example 1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"example 2", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"single match", []int{5}, 5, 0},
		{"single miss", []int{5}, 2, -1},
		{"first element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"last element", []int{1, 2, 3, 4, 5}, 5, 4},
		{"empty array", []int{}, 1, -1},
		{"two elements found right", []int{1, 4}, 4, 1},
		{"negative target", []int{-5, -3, -1, 2}, -3, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
