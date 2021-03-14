package main

func isAnagram(s string, t string) bool {
	var set = make(map[int32]int, 26)
	for _, i := range s {
		set[i] ++
	}
	for _, i := range t {
		set[i] --
	}
	for _, i := range set {
		if i != 0 {
			return false
		}
	}
	return true
}