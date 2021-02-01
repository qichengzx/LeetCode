package main

func main() {
	//fmt.Println()
}

func maximumWealth(accounts [][]int) int {
	var ans = 0
	for _, account := range accounts {
		account_wealth := 0
		for _, v := range account {
			account_wealth += v
		}
		if ans < account_wealth {
			ans = account_wealth
		}
	}
	return ans
}