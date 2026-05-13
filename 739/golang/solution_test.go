package problem739

import (
	"reflect"
	"testing"
)

// 739. Daily Temperatures
// Pattern: monotonic-stack

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		temp []int
		want []int
	}{
		{"example 1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"example 2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"example 3", []int{30, 60, 90}, []int{1, 1, 0}},
		{"decreasing", []int{90, 80, 70, 60}, []int{0, 0, 0, 0}},
		{"single", []int{50}, []int{0}},
		{"plateau", []int{50, 50, 50, 51}, []int{3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.temp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures(%v) = %v; want %v", tt.temp, got, tt.want)
			}
		})
	}
}
