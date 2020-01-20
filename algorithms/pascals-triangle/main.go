package main

func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		n := []int{}

		for inner := 0; inner <= i; inner++ {
			if inner == 0 || inner == i {
				n = append(n, 1)
			} else {
				n = append(n, res[i-1][inner-1]+res[i-1][inner])
			}
		}

		res = append(res, n)
	}
	return res
}
