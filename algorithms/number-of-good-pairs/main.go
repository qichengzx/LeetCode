package main

func numIdenticalPairs(nums []int) int {
	var ans = 0
	var repeat = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := repeat[nums[i]]; ok {
			if v == 1 {
				ans += 1
			} else {
				ans += v
			}
			repeat[nums[i]] += 1
		} else {
			repeat[nums[i]] = 1
		}
	}

	return ans
}
