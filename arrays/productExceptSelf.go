package arrays

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = 1
	}

	pre := 1
	for i := 0; i < len(nums); i++ {
		res[i] = pre
		pre *= nums[i]
	}

	post := 1
	for i := len(nums); i >= 0; i-- {
		res[i] *= post
		post *= nums[i]
	}

	return res
}
