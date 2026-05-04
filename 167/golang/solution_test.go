package problem167

import (
	"reflect"
	"testing"
)

// 167. Two Sum II
// Pattern: opposite-end

func TestSolution(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example 1: target at opposite ends",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "example 2: target in middle",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "example 3: negative numbers",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "two elements only",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "duplicates form the pair",
			numbers: []int{3, 3},
			target:  6,
			want:    []int{1, 2},
		},
		{
			name:    "duplicates with distinct pair elsewhere",
			numbers: []int{1, 2, 3, 3, 4},
			target:  7,
			want:    []int{3, 5},
		},
		{
			name:    "all negatives",
			numbers: []int{-5, -3, -1},
			target:  -4,
			want:    []int{1, 3},
		},
		{
			name:    "negative + positive",
			numbers: []int{-3, 1, 4, 9},
			target:  1,
			want:    []int{1, 3},
		},
		{
			name:    "first and last",
			numbers: []int{1, 3, 5, 7, 9},
			target:  10,
			want:    []int{1, 5},
		},
		{
			name:    "adjacent middle pair",
			numbers: []int{1, 2, 3, 4, 5, 6},
			target:  7,
			want:    []int{1, 6},
		},
		{
			name:    "includes zero",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "large gap",
			numbers: []int{1, 100, 200, 300},
			target:  301,
			want:    []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
