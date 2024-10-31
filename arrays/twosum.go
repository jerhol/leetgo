package twosum

func TwoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, num := range nums {
		if j, ok := lookup[target-num]; ok {
			return []int{j, i}
		}
		lookup[num] = i
	}
	return nil
}
