package tests

import (
	"reflect"
	"testing"

	twosum "leetgo/arrays"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{4, 5, 6, 8}, 11, []int{1, 2}},
		{[]int{2, 3, 5, 7}, 12, []int{2, 3}},
	}

	for _, tt := range tests {
		got := twosum.TwoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
