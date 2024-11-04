package twopointers

import (
	"strings"
	"unicode"
)

// Time: O(n)
// Space: O(1)

func IsPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	s = strings.Map(f, s)

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
