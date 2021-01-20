package main

func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	a := int32('a')
	m := make(map[string]struct{})
	for _, word := range words {
		s := ""
		for _, w := range word {
			s += codes[w-a]
		}

		m[s] = struct{}{}
	}

	return len(m)
}
