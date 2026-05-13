package problem20

import "testing"

// 20. Valid Parentheses
// Pattern: stack

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"example 1", "()", true},
		{"example 2", "()[]{}", true},
		{"example 3", "(]", false},
		{"nested", "([{}])", true},
		{"empty", "", true},
		{"unmatched open", "(", false},
		{"unmatched close", ")", false},
		{"wrong order", "([)]", false},
		{"deep nest", "((((()))))", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
