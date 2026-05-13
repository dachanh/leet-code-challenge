package problem102

import (
	"reflect"
	"testing"
)

// 102. Binary Tree Level Order Traversal
// Pattern: tree-bfs

func build(values []any) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []any
		want [][]int
	}{
		{"example 1", []any{3, 9, 20, nil, nil, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"example 2", []any{1}, [][]int{{1}}},
		{"empty", []any{}, [][]int{}},
		{"only left", []any{1, 2, nil, 3}, [][]int{{1}, {2}, {3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(build(tt.in))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
