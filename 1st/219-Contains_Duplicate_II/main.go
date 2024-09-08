package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	stock := make(map[int]int)

	for i, num := range nums {
		if j, exsits := stock[num]; exsits && i-j <= k {
			return true
		}
		stock[num] = i
	}
	return false

}
