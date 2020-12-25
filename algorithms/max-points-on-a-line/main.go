package main

import "math"

func maxPoints(points [][]int) int {
	size := len(points)
	if size == 0 {
		return 0
	}
	ans := 1
	for i := 0; i < size; i++ {
		if size-i <= ans {
			break
		}
		tmp, repeat := 1, 1
		for j := i + 1; j < size; j++ {
			if size-j+tmp <= ans {
				break
			}

			tmp++
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			if x1 == x2 && y1 == y2 {
				repeat++
				ans = int(math.Max(float64(ans), float64(tmp)))
				continue
			}

			for k := j + 1; k < size; k++ {
				if size-k+tmp <= ans {
					break
				}
				x3, y3 := points[k][0], points[k][1]
				if (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1) {
					tmp++
				}
			}

			ans = int(math.Max(float64(ans), float64(tmp)))
			tmp = repeat
		}
	}

	return ans
}
