package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok && i - m[v] <= k{
			return true
		} else {
			m[v] = i
		}
	}
	return false
}