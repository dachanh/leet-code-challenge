package problem21

import (
	"reflect"
	"testing"
)

// 21. Merge Two Sorted Lists
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
		a, b []int
		want []int
	}{
		{"example 1", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"both empty", []int{}, []int{}, []int{}},
		{"one empty", []int{}, []int{0}, []int{0}},
		{"disjoint", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"interleaved", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"duplicates", []int{1, 1, 1}, []int{1, 1}, []int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(mergeTwoLists(build(tt.a), build(tt.b)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
