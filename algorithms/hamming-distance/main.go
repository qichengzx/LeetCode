package main

func hammingDistance(x int, y int) int {
	tmp := x^y
	dis := 0

	for tmp > 0 {
		if (tmp >> 1) << 1 != tmp {
			dis ++
		}

		tmp >>= 1
	}

	return dis
}
