package problem567

import "testing"

// 567. Permutation in String
// Pattern: sliding-window
// Signature: func checkInclusion(s1 string, s2 string) bool

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "example 1: ba is a permutation of ab",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "example 2: no permutation",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
		{
			name: "s1 equals s2",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "s1 longer than s2",
			s1:   "abcd",
			s2:   "abc",
			want: false,
		},
		{
			name: "single char present",
			s1:   "a",
			s2:   "ba",
			want: true,
		},
		{
			name: "single char absent",
			s1:   "a",
			s2:   "bcdef",
			want: false,
		},
		{
			name: "duplicates require correct multiplicity",
			s1:   "aab",
			s2:   "eidbaaooo",
			want: true,
		},
		{
			name: "duplicates with wrong multiplicity",
			s1:   "aab",
			s2:   "eidabooo",
			want: false,
		},
		{
			name: "permutation at the start",
			s1:   "abc",
			s2:   "cbaxyz",
			want: true,
		},
		{
			name: "permutation at the very end",
			s1:   "abc",
			s2:   "xyzcba",
			want: true,
		},
		{
			name: "all same letter",
			s1:   "aaa",
			s2:   "aaaa",
			want: true,
		},
		{
			name: "no overlap of letters",
			s1:   "xyz",
			s2:   "abcdefg",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v; want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
