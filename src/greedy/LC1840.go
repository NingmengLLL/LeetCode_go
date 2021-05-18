package main

import (
	"fmt"
	"sort"
)

type ArraySlice [][]int

func (s ArraySlice) Len() int {
	return len(s)
}

func (s ArraySlice) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s ArraySlice) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxBuilding(n int, r [][]int) int {

	r = append(r, []int{1,0})
	sort.Sort(ArraySlice(r))
	if r[len(r)-1][0] !=n {
		r = append(r, []int{n, n-1})
	}

	length := len(r)
	for i := 1; i < length; i++ {
		r[i][1] = min(r[i][1], r[i - 1][1] + (r[i][0] - r[i - 1][0]))
	}
	for i := length - 2; i >= 0; i-- {
		r[i][1] = min(r[i][1], r[i + 1][1] + (r[i + 1][0] - r[i][0]))
	}
	ans := 0
	for i := 0; i < length - 1; i++ {
		best := ((r[i + 1][0] - r[i][0]) + r[i][1] + r[i + 1][1]) / 2
		ans = max(ans, best)
	}
	return ans
}

func main() {
	//fmt.Println(maxBuilding(6,nil))
	s := []int{5}

	s = append(s,7)
	fmt.Println(s)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])
	//
	//s = append(s,9)
	//fmt.Println(s)
	//fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0])

	x := append(s, 11)
	fmt.Println(x)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(x) =", &x[0])
	x[0] = -1

	y := append(s, 12)
	fmt.Println(y)
	fmt.Println("cap(s) =", cap(s), "ptr(s) =", &s[0], "ptr(y) =", &y[0])
	fmt.Println(x)
}
