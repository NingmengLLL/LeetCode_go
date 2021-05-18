package main

func countNicePairs(nums []int) int {
	m:=make(map[int]int,0)
	ans:=0
	mod:=1000000007
	for _,num:=range nums{
		k:=num-rec(num)
		ans=(ans+m[k])%mod
		m[k]++
	}
	return ans
}

func rec(n int) int {
	m:=0
	for n>0{
		k:=n%10
		n = n/10
		m = m*10+k
	}
	return m
}

