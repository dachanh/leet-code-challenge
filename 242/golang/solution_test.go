package problem242

import "testing"

// 242. Valid Anagram
// Pattern: hashmap

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"example 1", "anagram", "nagaram", true},
		{"example 2", "rat", "car", false},
		{"different lengths", "ab", "abc", false},
		{"both empty", "", "", true},
		{"single same", "a", "a", true},
		{"single diff", "a", "b", false},
		{"duplicates", "aabbcc", "abcabc", true},
		{"same letters diff counts", "aab", "abb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v; want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
