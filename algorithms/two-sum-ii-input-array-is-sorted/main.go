package main

func twoSum(numbers []int, target int) []int {
	var m = make(map[int]int, len(numbers))

	for i, v := range numbers {
		if m[target-v] != 0 {
			return []int{m[target-v], i + 1}
		}

		m[v] = i + 1
	}

	return nil
}
