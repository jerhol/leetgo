package tests

import (
	"leetgo/arrays"
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{[]int{0, 1, 2, 3, 4, 3}, true},
		{[]int{0, 1, 2, 3, 4, 5}, false},
	}

	for _, tt := range tests {
		got := arrays.HasDuplicate(tt.input)
		if got != tt.want {
			t.Errorf("HasDuplicate(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
