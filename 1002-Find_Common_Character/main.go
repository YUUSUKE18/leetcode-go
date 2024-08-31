package findcommoncharacter

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	charCount := make(map[rune]int)

	for _, char := range words[0] {
		charCount[char]++
	}

	for _, word := range words[1:] {
		currentCount := make(map[rune]int)
		for _, w := range word {
			currentCount[w]++
		}

		for c := range charCount {
			if currentCount[c] == 0 {
				delete(charCount, c)
			}
		}
	}

	var results []string
	for char, count := range charCount {
		for i := 0; i < count; i++ {
			results = append(results, string(char))
		}
	}

	return results
}
