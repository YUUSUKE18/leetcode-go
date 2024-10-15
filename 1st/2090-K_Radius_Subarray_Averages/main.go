package kradiussubarrayaverages

func getAverages(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = -1
	}

	if 2*k+1 > n {
		return result
	}

	beforeSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		beforeSum[i+1] = beforeSum[i] + int64(nums[i])
	}

	for i := k; i < n-k; i++ {
		sum := beforeSum[i+k+1] - beforeSum[i-k]
		result[i] = int(sum / int64(2*k+1))
	}
	return result
}
