package main

func permute(nums []int) [][]int {
	var res [][]int

	swap := func (i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}

		for i := index; i < len(nums); i++ {
			swap(index, i)
			backtrack(index + 1)
			swap(index, i)
		}
	}
 	backtrack(0)
	return res
}
