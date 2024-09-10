package crawlerlogfolder

func minOperations(logs []string) int {
	depth := 0

	for _, log := range logs {
		switch log {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":

		default:
			depth++
		}
	}

	return depth
}
