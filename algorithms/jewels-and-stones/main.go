package main

func numJewelsInStones(J string, S string) int {
	num := 0
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(J); j++ {
			if J[j] == S[i] {
				num++
			}
		}
	}

	return num
}
