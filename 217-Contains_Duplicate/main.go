package containsduplicate

func containsDuplicate(nums []int) bool {
	stock := make(map[int]int)

	for _, num := range nums {
		if stock[num] == num {
			return true
		}
		stock[num] = num
	}
	return false

}
