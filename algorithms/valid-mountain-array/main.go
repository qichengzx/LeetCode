package main

func validMountainArray(A []int) bool {
	l := len(A)
	if l < 3 {
		return false
	}

	i := 0
	for i+1 < l && A[i] < A[i+1] {
		i++
	}

	if i == 0 || i == l-1 {
		return false
	}

	for i+1 < l && A[i] > A[i+1] {
		i++
	}

	return i == l-1
}
