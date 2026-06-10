package solution

import (
	"math"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"example", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"single", []int{-1}, 1, -1},
		{"allNegative", []int{-5, -2, -3, -4}, 2, -2.5},
		{"wholeArray", []int{5, 6, 7}, 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.nums, tt.k)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Fatalf("findMaxAverage(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
