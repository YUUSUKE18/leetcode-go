package customsortstring

func customSortString(order string, s string) string {
	orderMap := make(map[rune]int)

	// orderをハッシュマップに挿入する
	for i, o := range order {
		orderMap[o] = i
	}
	// sの文字をカウント
	sCount := make(map[rune]int)
	for _, ch := range s {
		sCount[ch]++
	}
	// 結果を格納する
	result := make([]rune, 0, len(s))
	for _, ch := range order {
		for i := 0; i < sCount[ch]; i++ {
			result = append(result, ch)
		}
		delete(sCount, ch)
	}

	for ch, count := range sCount {
		for i := 0; i < count; i++ {
			result = append(result, ch)
		}
	}

	return string(result)
}
