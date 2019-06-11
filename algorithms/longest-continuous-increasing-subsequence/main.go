package main

import "math"

func findLengthOfLCIS(nums []int) int {
	max, count := float64(0),float64(0)
	for i:= 0; i< len(nums); i++ {
		if i-1 >= 0 && nums[i] <= nums[i-1] {
			max = math.Max(max, count)
			count = 0
		}
		count ++
	}
	max = math.Max(max, count)
	return int(max)
}

