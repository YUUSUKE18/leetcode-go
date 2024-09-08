package removeelement

func removeElement(nums []int, val int) int {

	// numsが0の場合、0を返す
	if len(nums) == 0 {
		return 0
	}
	// カウント設定
	var c = 0
	// ループ処理にて、ターゲットと一致しないものを見つけると
	// 位置を移動して、cをカウントアップする。
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[c] = nums[i]
			c++
		}
	}
	// 最終処理を戻す
	return c
}
