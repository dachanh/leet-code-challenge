package problem76

import "testing"

// 76. Minimum Window Substring
// Pattern: sliding-window
// Signature: func minWindow(s string, t string) string

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1: classic LeetCode case",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2: exact match",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "example 3: t longer than s",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "no characters of t in s",
			s:    "abcdef",
			t:    "xyz",
			want: "",
		},
		{
			name: "duplicates in t require multiplicity",
			s:    "aabdec",
			t:    "abc",
			want: "abdec",
		},
		{
			name: "duplicates in t and s",
			s:    "aaflslflsldkalskaaa",
			t:    "aaa",
			want: "aaa",
		},
		{
			name: "t equals s",
			s:    "abc",
			t:    "abc",
			want: "abc",
		},
		{
			name: "empty s",
			s:    "",
			t:    "a",
			want: "",
		},
		{
			name: "empty t",
			s:    "abc",
			t:    "",
			want: "",
		},
		{
			name: "case-sensitive: uppercase vs lowercase distinct",
			s:    "aA",
			t:    "Aa",
			want: "aA",
		},
		{
			name: "window at the very end",
			s:    "xyzABC",
			t:    "ABC",
			want: "ABC",
		},
		{
			name: "window at the very start",
			s:    "ABCxyz",
			t:    "ABC",
			want: "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q; want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
