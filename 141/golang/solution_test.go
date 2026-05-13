package problem141

import "testing"

// 141. Linked List Cycle
// Pattern: fast-slow

func buildList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
		want   bool
	}{
		{"example 1", []int{3, 2, 0, -4}, 1, true},
		{"example 2", []int{1, 2}, 0, true},
		{"example 3", []int{1}, -1, false},
		{"empty", []int{}, -1, false},
		{"single no cycle", []int{1}, -1, false},
		{"single self cycle", []int{1}, 0, true},
		{"no cycle", []int{1, 2, 3, 4, 5}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.values, tt.pos)
			if got := hasCycle(head); got != tt.want {
				t.Errorf("hasCycle(%v, pos=%d) = %v; want %v", tt.values, tt.pos, got, tt.want)
			}
		})
	}
}
