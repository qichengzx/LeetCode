package main

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}

	index := 0
	for _, v := range s {
		if m[v-'a'] == 1 {
			return index
		} else {
			index++
		}
	}

	return -1
}