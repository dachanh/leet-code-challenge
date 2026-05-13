package problem226

import (
	"reflect"
	"testing"
)

// 226. Invert Binary Tree
// Pattern: tree-dfs

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

func levelOrder(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}
	out := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []any
		want []any
	}{
		{"example 1", []any{4, 2, 7, 1, 3, 6, 9}, []any{4, 7, 2, 9, 6, 3, 1}},
		{"example 2", []any{2, 1, 3}, []any{2, 3, 1}},
		{"empty", []any{}, []any{}},
		{"single", []any{1}, []any{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(invertTree(build(tt.in)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
