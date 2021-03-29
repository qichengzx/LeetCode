package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	pos, cnt := 0, 0
	if len(items) == 0 {
		return cnt
	}
	if ruleKey == "type" {
		pos = 0
	} else if ruleKey == "color" {
		pos = 1
	} else {
		pos = 2
	}
	for _, item := range items {
		if item[pos] == ruleValue {
			cnt++
		}
	}
	return cnt
}
