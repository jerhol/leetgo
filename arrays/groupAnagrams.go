package arrays

import "fmt"

// Time: O(m * n)
// Space: O(m)
func GroupAnagrams(strs []string) [][]string {
	sMap := make(map[string][]string)

	for _, str := range strs {
		count := make([]int, 26)
		for _, char := range str {
			count[char-'a']++
		}

		key := fmt.Sprintf("%v", count)

		sMap[key] = append(sMap[key], str)
	}

	var res [][]string
	for _, group := range sMap {
		res = append(res, group)
	}

	return res
}
