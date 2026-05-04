package problem904

import "testing"

// 904. Fruit Into Baskets
// Pattern: sliding-window
// Signature: func totalFruit(fruits []int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "example 1: pick whole array",
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "example 2: drop first 0 to get [1,2,1,2]",
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
		{
			name:   "example 3: longest two-fruit window in middle",
			fruits: []int{1, 2, 3, 2, 2},
			want:   4,
		},
		{
			name:   "single tree",
			fruits: []int{5},
			want:   1,
		},
		{
			name:   "two trees, same fruit",
			fruits: []int{3, 3},
			want:   2,
		},
		{
			name:   "two trees, different fruits",
			fruits: []int{3, 4},
			want:   2,
		},
		{
			name:   "all same fruit",
			fruits: []int{7, 7, 7, 7, 7},
			want:   5,
		},
		{
			name:   "all distinct fruits",
			fruits: []int{1, 2, 3, 4, 5},
			want:   2,
		},
		{
			name:   "long window then break",
			fruits: []int{1, 1, 1, 2, 2, 3, 1, 1},
			want:   5,
		},
		{
			name:   "tail window is the longest",
			fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := totalFruit(tt.fruits)
			if got != tt.want {
				t.Errorf("totalFruit(%v) = %d; want %d", tt.fruits, got, tt.want)
			}
		})
	}
}
