package twosum

func twoSum(nums []int, target int) []int {
	// 数値を格納するハッシュマップ
	combination := make(map[int]int)
	for i, num := range nums {
		// targetとの差分確認
		complement := target - num
		if j, found := combination[complement]; found {
			return []int{i, j}
		}
		combination[num] = i
	}
	return []int{}
}
