package main

func longestBeautifulSubstring(word string) int {

	res := 0
	curRes := 1
	count := 1

	for i:=1;i<len(word);i++ {
		if word[i]>=word[i-1] {
			curRes++
		}
		if word[i]>word[i-1] {
			count++
		}
		if word[i] <word[i-1]{
			curRes=1
			count=1
		}
		if count==5 && curRes>res {
			res = curRes
		}
	}

	return res
}


