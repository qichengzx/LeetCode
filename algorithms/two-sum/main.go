package main

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := m[target-v]; ok{
			return []int{j, i}
		}

		m[v] = i
	}

	return nil
}
