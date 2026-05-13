package problem49

import (
	"sort"
	"testing"
)

// 49. Group Anagrams
// Pattern: hashmap

func normalize(groups [][]string) [][]string {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func equal(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			"example 1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{"empty string", []string{""}, [][]string{{""}}},
		{"single", []string{"a"}, [][]string{{"a"}}},
		{"all same", []string{"abc", "bca", "cab"}, [][]string{{"abc", "bca", "cab"}}},
		{"no anagrams", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalize(groupAnagrams(tt.strs))
			want := normalize(tt.want)
			if !equal(got, want) {
				t.Errorf("groupAnagrams(%v) = %v; want %v", tt.strs, got, want)
			}
		})
	}
}
