package convertbinarynumberinalinkedlisttointeger

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	// 戻り値
	result := 0

	for head != nil {
		result = result*2 + head.Val
		head = head.Next
	}
	return result
}
