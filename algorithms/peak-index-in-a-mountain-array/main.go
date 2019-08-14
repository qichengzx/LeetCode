package main

func peakIndexInMountainArray(A []int) int {
	i := 0
	for A[i] < A[i+1] {
		i++
	}
	return i
}
