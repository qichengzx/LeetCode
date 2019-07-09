package main

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := 0
	l := len(s)
	for i := 0;i < l; i++{
		if i < l-1 {
			if m[rune(s[i])] < m[rune(s[i+1])] {
				n -= m[rune(s[i])]
			} else {
				n += m[rune(s[i])]
			}
		}
		if i == l -1{
			n += m[rune(s[i])]
		}

	}

	return n
}
