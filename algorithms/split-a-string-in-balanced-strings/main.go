package main

func balancedStringSplit(s string) int {
	count, num := 0, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "L" {
			num += 1
		} else {
			num -= 1
		}
		if num == 0 {
			count++
		}
	}

	return count
}
