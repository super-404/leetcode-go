package solution

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"example", "ADOBECODEBANC", "ABC", "BANC"},
		{"singleMatch", "a", "a", "a"},
		{"singleMissing", "a", "aa", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
