package arrays

// Time: O(n)
// Space: O(1)

func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, char := range count {
		if char != 0 {
			return false
		}
	}

	return true

}
