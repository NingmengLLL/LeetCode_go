package main

import "fmt"

func minimumSwap(s1 string, s2 string) int {
	
	count1 := 0
	count2 := 0
	xy := "xy"
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if s1[i] == xy[0] {
				count1++
			} else {
				count2++
			}
		}
	}
	if (count1 + count2) % 2 != 0 {
		return -1
	}
	res := count1/2 + count2/2
	if count1 % 2 !=0 {
		res += 2
	}
	return res
}

func main() {
	 c := 'x'
	 s := "xy"
	 fmt.Printf("%T, %T", c, s[0])
	 fmt.Println('x'==s[0])
}
