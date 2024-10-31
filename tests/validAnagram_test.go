package tests

import (
	"leetgo/arrays"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"racecar", "rraceca", true},
		{"cat", "bat", false},
	}

	for _, tt := range tests {
		got := arrays.ValidAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("ValidAnagram(%v, %v) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
