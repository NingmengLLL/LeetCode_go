package main

import (
	"sort"
)

func maxFrequency(nums []int, k int) int {

	cost := 0
	sort.Ints(nums)
	l, r := 0, 1
	res := 1
	for r < len(nums) {
		cost += (nums[r] - nums[r-1]) * (r-l)
		for cost > k {
			cost -= nums[r] - nums[l]
			l++
		}
		if r-l+1 > res {
			res = r-l+1
		}
	}
	return res
}

func main() {

}
