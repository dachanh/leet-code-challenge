package problem200

import "testing"

// 200. Number of Islands
// Pattern: graph-dfs

func toGrid(strs []string) [][]byte {
	g := make([][]byte, len(strs))
	for i, s := range strs {
		g[i] = []byte(s)
	}
	return g
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want int
	}{
		{
			"example 1",
			[]string{
				"11110",
				"11010",
				"11000",
				"00000",
			}, 1,
		},
		{
			"example 2",
			[]string{
				"11000",
				"11000",
				"00100",
				"00011",
			}, 3,
		},
		{"all water", []string{"000", "000"}, 0},
		{"single land", []string{"1"}, 1},
		{"diagonal not connected", []string{"101", "010", "101"}, 5},
		{"row of islands", []string{"10101"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(toGrid(tt.grid)); got != tt.want {
				t.Errorf("numIslands(%v) = %d; want %d", tt.grid, got, tt.want)
			}
		})
	}
}
