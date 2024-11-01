package arrays

// Time: O(m * n)
// Space: O(m)
func GroupAnagrams(strs []string) [][]string {
	sMap := map[[26]int][]string{}

	for _, str := range strs {
		count := [26]int{}
		for _, char := range str {
			count[char-'a'] += 1
		}

		sMap[count] = append(sMap[count], str)
	}

	res := [][]string{}
	for _, group := range sMap {
		res = append(res, group)
	}

	return res
}
