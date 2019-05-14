package main

func flipAndInvertImage(A [][]int) [][]int {
	l := len(A[0])
	for _, row :=  range A {
		for i := 0; i < (l+1)/2 ; i++ {
			k := l - 1 - i
			row[i], row[k] = row[k]^1, row[i] ^ 1
		}
	}

	return A
}
