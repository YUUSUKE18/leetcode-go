package maximumaveragesubarrayi

func findMaxAverage(nums []int, k int) float64 {
	// nums = 0の場合、0を戻す
	if len(nums) == 0 {
		return 0
	}
	// nums = 0の場合、0を戻す
	if len(nums) == 0 {
		return 0
	}
	// 配列[0-3]までの合計
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	// k番目以降len(nums)までの合計
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}

	}
	// 戻す
	return float64(maxSum) / float64(k)
}
