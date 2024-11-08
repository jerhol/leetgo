package twopointers

// Time: O(n)
// Space: O(1) *

func TwoSumSorted(nums []int, t int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]+nums[j] == t {
			return []int{i + 1, j + 1}
		}
		if nums[i]+nums[j] > t {
			j--
		} else {
			i++
		}
	}

	return []int{}
}
