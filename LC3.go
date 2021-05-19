package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var res int
	var subStr string
	left := 0
	right := 0
	subStr = s[left:right]
	for right < len(s) {

		if index := strings.IndexByte(subStr, s[right]); index != -1 {
			left += index + 1
		}
		subStr = s[left:right+1]
		if len(subStr) > res {
			res = len(subStr)
		}
	}
	return res
}

func main(){
	s := make([]string, 10)
	for i := 0; i < 3; i++ {
		s[i] = "tcy"
	}
	fmt.Println(strings.Join(s, "最棒"))//tcy最棒tcy最棒tcy最棒最棒最棒最棒最棒最棒最棒

}
