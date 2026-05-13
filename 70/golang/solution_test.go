package problem70

import "testing"

// 70. Climbing Stairs
// Pattern: dp

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=1", 1, 1},
		{"n=2", 2, 2},
		{"n=3", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=10", 10, 89},
		{"n=20", 20, 10946},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
