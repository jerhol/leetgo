package tests

import (
	"leetgo/arrays"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"act", "pots", "tops", "cat", "stop", "hat"},
			[][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
	}

	for _, tt := range tests {
		got := arrays.GroupAnagrams(tt.strs)

		for _, group := range got {
			sort.Strings(group)
		}
		for _, group := range tt.want {
			sort.Strings(group)
		}

		sort.Slice(got, func(i, j int) bool {
			return got[i][0] < got[j][0]
		})
		sort.Slice(tt.want, func(i, j int) bool {
			return tt.want[i][0] < tt.want[j][0]
		})

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GroupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}
