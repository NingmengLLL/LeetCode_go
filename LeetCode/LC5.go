package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := helper(s, i, i)
		left2, right2 := helper(s, i, i+1)
		if right1-left1 > right-left {
			left, right = left1, right1
		}
		if right2-left2 > right-left {
			left, right = left2, right2
		}
	}
	return s[left:right+1]
}

func helper(s string, left int, right int)(int, int){
	for left >=0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
