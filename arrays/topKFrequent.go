package arrays

// Time: O(n)
// Space: O(n)

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	res := []int{}

	for _, num := range nums {
		count[num]++
	}

	for num, freqCount := range count {
		freq[freqCount] = append(freq[freqCount], num)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
