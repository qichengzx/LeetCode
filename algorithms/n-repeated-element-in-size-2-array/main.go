package main

func repeatedNTimes(A []int) int {
	m := make(map[int]struct{}, len(A))
	for _, v := range A {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}

	return 0
}
