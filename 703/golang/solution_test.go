package problem703

import "testing"

// 703. Kth Largest Element in a Stream
// Pattern: heap

func TestSolution(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		kl := Constructor(3, []int{4, 5, 8, 2})
		ops := []struct{ val, want int }{
			{3, 4},
			{5, 5},
			{10, 5},
			{9, 8},
			{4, 8},
		}
		for _, op := range ops {
			if got := kl.Add(op.val); got != op.want {
				t.Errorf("Add(%d) = %d; want %d", op.val, got, op.want)
			}
		}
	})

	t.Run("k=1 always largest", func(t *testing.T) {
		kl := Constructor(1, []int{})
		ops := []struct{ val, want int }{
			{5, 5}, {3, 5}, {10, 10}, {7, 10},
		}
		for _, op := range ops {
			if got := kl.Add(op.val); got != op.want {
				t.Errorf("Add(%d) = %d; want %d", op.val, got, op.want)
			}
		}
	})

	t.Run("with negatives", func(t *testing.T) {
		kl := Constructor(2, []int{-1, -5})
		if got := kl.Add(-3); got != -3 {
			t.Errorf("Add(-3) = %d; want -3", got)
		}
		if got := kl.Add(0); got != -1 {
			t.Errorf("Add(0) = %d; want -1", got)
		}
	})
}
