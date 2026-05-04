package problem11

import "testing"

// 11. Container With Most Water
// Pattern: opposite-end
// Signature: func maxArea(height []int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1: classic LeetCode case",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "example 2: two bars",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "ascending heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "descending heights",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "all equal heights",
			height: []int{4, 4, 4, 4, 4},
			want:   16,
		},
		{
			name:   "tall walls at the ends",
			height: []int{8, 1, 1, 1, 1, 1, 8},
			want:   48,
		},
		{
			name:   "single tall in middle",
			height: []int{1, 2, 1, 10, 1, 2, 1},
			want:   8,
		},
		{
			name:   "zeros mixed in",
			height: []int{0, 2, 0, 4, 0, 5},
			want:   8,
		},
		{
			name:   "two-element minimum",
			height: []int{2, 3},
			want:   2,
		},
		{
			name:   "wide low container beats narrow tall",
			height: []int{2, 3, 4, 5, 18, 17, 6},
			want:   17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
