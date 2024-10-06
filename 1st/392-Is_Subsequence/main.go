package issubsequence

func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}

	return j == len(s)
}
