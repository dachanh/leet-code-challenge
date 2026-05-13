package problem75

import (
	"reflect"
	"testing"
)

// 75. Sort Colors
// Pattern: dutch-national-flag

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example 1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"example 2", []int{2, 0, 1}, []int{0, 1, 2}},
		{"single 0", []int{0}, []int{0}},
		{"single 1", []int{1}, []int{1}},
		{"single 2", []int{2}, []int{2}},
		{"all zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"already sorted", []int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"reverse", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.in...)
			sortColors(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("sortColors(%v) -> %v; want %v", tt.in, nums, tt.want)
			}
		})
	}
}
