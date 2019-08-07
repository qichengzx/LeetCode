package main

func toLowerCase(str string) string {
	runeNew := []rune{}
	for _, r := range []rune(str) {
		if r >= 65 && r <= 90 {
			runeNew = append(runeNew, r+32)
		} else {
			runeNew = append(runeNew, r)
		}
	}

	return string(runeNew)
}
