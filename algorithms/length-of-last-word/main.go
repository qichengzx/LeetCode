package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	ans := 0
	s = strings.Trim(s, " ")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			ans += 1
		} else {
			return ans
		}
	}
	return ans
}
