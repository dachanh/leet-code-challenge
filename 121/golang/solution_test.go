package problem121

import "testing"

// 121. Best Time to Buy and Sell Stock
// Pattern: greedy

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example 2", []int{7, 6, 4, 3, 1}, 0},
		{"empty", []int{}, 0},
		{"single", []int{5}, 0},
		{"two ascending", []int{1, 5}, 4},
		{"two descending", []int{5, 1}, 0},
		{"valley", []int{3, 2, 6, 5, 0, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
			}
		})
	}
}
