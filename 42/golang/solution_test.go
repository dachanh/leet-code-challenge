package problem42

import "testing"

// 42. Trapping Rain Water
// Pattern: opposite-end
// Signature: func trap(height []int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1: classic LeetCode case",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "example 2: simple valley",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "empty input",
			height: []int{},
			want:   0,
		},
		{
			name:   "single bar",
			height: []int{5},
			want:   0,
		},
		{
			name:   "two bars cannot trap",
			height: []int{2, 5},
			want:   0,
		},
		{
			name:   "monotonically increasing",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "monotonically decreasing",
			height: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "flat surface",
			height: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "single deep valley",
			height: []int{5, 0, 5},
			want:   5,
		},
		{
			name:   "asymmetric walls",
			height: []int{3, 0, 0, 0, 5},
			want:   9,
		},
		{
			name:   "plateau in the middle",
			height: []int{4, 1, 1, 1, 4},
			want:   9,
		},
		{
			name:   "all zeros",
			height: []int{0, 0, 0, 0},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("trap(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
