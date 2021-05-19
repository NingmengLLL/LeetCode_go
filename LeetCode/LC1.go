package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	indexMap := map[int]int {}
	for i, x := range nums {
		if v, ok := indexMap[target-x]; ok {
			return []int{v, i}
		}
		indexMap[x] = i
	}
	return nil
}

func main()  {
	a := []int{1,2,3}
	b := 5
	fmt.Println(twoSum(a,b))
}
