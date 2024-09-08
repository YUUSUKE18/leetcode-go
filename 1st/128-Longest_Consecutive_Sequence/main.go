package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	numHash := make(map[int]bool)
	for _, num := range nums {
		numHash[num] = true
	}
	longest := 0
	currnetNum := 0
	for num := range numHash {
		if !numHash[num-1] {
			currnetNum = num
			currentStreak := 1
			for numHash[currnetNum+1] {
				currnetNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
