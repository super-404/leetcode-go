package solution

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"example1", "abcabcbb", 3},
		{"example2", "bbbbb", 1},
		{"example3", "pwwkew", 3},
		{"duplicateInsideWindow", "abba", 2},
		{"withSpace", "a b c a", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Fatalf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
