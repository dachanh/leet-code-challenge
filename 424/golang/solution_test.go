package problem424

import "testing"

// 424. Longest Repeating Character Replacement
// Pattern: sliding-window
// Signature: func characterReplacement(s string, k int) int

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "example 1: replace 1 to extend AB to AB+B",
			s:    "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "example 2: replace one A to get BBBB",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "single character",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "all same characters with no replacement",
			s:    "AAAA",
			k:    0,
			want: 4,
		},
		{
			name: "all distinct, k zero",
			s:    "ABCDE",
			k:    0,
			want: 1,
		},
		{
			name: "all distinct, enough k to convert all",
			s:    "ABCDE",
			k:    4,
			want: 5,
		},
		{
			name: "k larger than string length",
			s:    "ABC",
			k:    10,
			want: 3,
		},
		{
			name: "k zero with one repeating run",
			s:    "AABBBCC",
			k:    0,
			want: 3,
		},
		{
			name: "alternating with k=1",
			s:    "ABABAB",
			k:    1,
			want: 3,
		},
		{
			name: "long run dominated",
			s:    "AAAABBBBCCCC",
			k:    2,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
