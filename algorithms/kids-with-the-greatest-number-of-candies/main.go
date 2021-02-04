package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var ans = make([]bool, len(candies))
	greatest := candies[0]
	for i := 0; i < len(candies); i++ {
		if candies[i] > greatest {
			greatest = candies[i]
		}
	}
	for i, candy := range candies {
		if candy+extraCandies >= greatest {
			ans[i] = true
		} else {
			ans[i] = false
		}
	}
	return ans
}
