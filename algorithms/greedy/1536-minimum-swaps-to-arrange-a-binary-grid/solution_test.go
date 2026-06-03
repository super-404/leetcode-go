package solution

import (
	"testing"
)

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}, 3},
		{"示例2-不可能", [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}, -1},
		{"示例3-已有效", [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSwaps(tt.grid)
			if got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
