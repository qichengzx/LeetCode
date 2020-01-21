package main

import "strings"

func intToRoman(num int) string {
	var v = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var s = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var b strings.Builder
	for i := 0; i < len(v); i++ {
		for num >= v[i] {
			num -= v[i]
			b.WriteString(s[i])
		}
	}

	return b.String()
}
