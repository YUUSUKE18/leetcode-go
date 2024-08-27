package validpalindrome

func isPalindrome(s string) bool {
	var l = 0
	var r = len(s) - 1

	for l < r {
		for l < r && !alphaNum(s[l]) {
			l++
		}
		for r > l && !alphaNum(s[r]) {
			r--
		}
		if lower(s[l]) != lower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

// ASCIIに適合するか確認する。
func alphaNum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

// 大文字から小文字にする処理
func lower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
