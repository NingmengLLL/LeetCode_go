package main

import (
	"strings"
	"unicode"
)

func numDifferentIntegers(word string) int {

	set := map[string]struct{}{}
	for _, s := range strings.FieldsFunc(word, unicode.IsLower){
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		set[s] = struct{}{}
	}
	return len(set)
}
