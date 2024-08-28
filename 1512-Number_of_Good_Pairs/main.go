package numberofgoodpairs

func numIdenticalPairs(nums []int) int {
	// hashMapの作成
	count := make(map[int]int)
	c := 0
	// ループ処理
	for _, num := range nums {
		if prevCount, exists := count[num]; exists {
			c += prevCount
		}
		count[num]++
	}

	return c
}
