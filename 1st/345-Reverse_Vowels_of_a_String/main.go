package reversevowelsofastring

import "strings"

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !strings.ContainsRune(vowels, runes[left]) {
			left++
		}
		for left < right && !strings.ContainsRune(vowels, runes[right]) {
			right--
		}

		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
