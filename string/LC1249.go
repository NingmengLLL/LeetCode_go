package main

import "strings"

func minRemoveToMakeValid(s string) string {
	sum := 0
	length := len(s)
	tmp := []rune(s)
	for i:=0;i< length;i++{
		if tmp[i] == '('{
			sum++
		}else if tmp[i] == ')'{
			sum--
			if sum < 0{
				tmp[i] = '0'
				sum = 0
			}
		}
	}
	sum = 0
	for i:= length -1;i>=0;i--{
		if tmp[i] == ')'{
			sum++
		}else if tmp[i] == '('{
			sum--
			if sum < 0{
				tmp[i] = '0'
				sum = 0
			}
		}
	}
	return strings.ReplaceAll(string(tmp),"0","")
}

