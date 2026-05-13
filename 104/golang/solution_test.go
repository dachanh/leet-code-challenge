package problem104

import "testing"

// 104. Maximum Depth of Binary Tree
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

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []any
		want int
	}{
		{"example 1", []any{3, 9, 20, nil, nil, 15, 7}, 3},
		{"example 2", []any{1, nil, 2}, 2},
		{"empty", []any{}, 0},
		{"single", []any{1}, 1},
		{"left chain", []any{1, 2, nil, 3, nil, 4}, 4},
		{"balanced", []any{1, 2, 3, 4, 5, 6, 7}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(build(tt.in)); got != tt.want {
				t.Errorf("maxDepth(%v) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}
