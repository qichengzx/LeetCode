package main

func restoreString(s string, indices []int) string {
	var ns = make([]uint8, len(indices))
	for i, index := range indices {
		ns[index] = s[i]
	}
	return string(ns)
}