package problem206

import (
	"reflect"
	"testing"
)

// 206. Reverse Linked List
// Pattern: linked-list

func build(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	curr := head
	for _, v := range values[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func toSlice(head *ListNode) []int {
	out := []int{}
	for curr := head; curr != nil; curr = curr.Next {
		out = append(out, curr.Val)
	}
	return out
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"example 1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"example 2", []int{1, 2}, []int{2, 1}},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"duplicates", []int{1, 1, 2, 2}, []int{2, 2, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(reverseList(build(tt.in)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
