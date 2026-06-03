package solution

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"示例1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"示例2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"示例3", []int{9}, []int{1, 0}},
		{"多位9", []int{9, 9}, []int{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
