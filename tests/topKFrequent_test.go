package tests

import (
	"leetgo/arrays"
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 2, []int{3, 2}},
		{[]int{7, 7}, 1, []int{7}},
	}

	for _, tt := range tests {
		got := arrays.TopKFrequent(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TopKFrequent(%v, %d) = %v; want %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
