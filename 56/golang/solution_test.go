package problem56

import (
	"reflect"
	"testing"
)

// 56. Merge Intervals
// Pattern: sort

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want [][]int
	}{
		{"example 1", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"example 2", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"single", [][]int{{1, 4}}, [][]int{{1, 4}}},
		{"contained", [][]int{{1, 10}, {2, 3}, {4, 5}}, [][]int{{1, 10}}},
		{"unsorted", [][]int{{5, 8}, {1, 3}, {6, 9}}, [][]int{{1, 3}, {5, 9}}},
		{"touch", [][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 4}}},
		{"disjoint", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
