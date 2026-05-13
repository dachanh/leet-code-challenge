package problem1

import (
	"reflect"
	"testing"
)

// 1. Two Sum
// Pattern: hashmap

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
		{"negatives", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"with zero", []int{0, 4, 3, 0}, 0, []int{0, 3}},
		{"first and last", []int{1, 5, 7, 2, 9}, 10, []int{0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
