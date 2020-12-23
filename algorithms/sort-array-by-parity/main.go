package main

func sortArrayByParity(A []int) []int {
	var odd []int
	var even []int
	l := len(A)
	for i := 0; i< l; i++{
		if A[i] % 2 == 0 {
			even = append(even, A[i])
		} else {
			odd = append(odd, A[i])
		}
	}

	return append(even, odd...)
}
