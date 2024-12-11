package removelinkedlistelements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// output Node
	result := &ListNode{0, head}
	count := result

	for count.Next != nil {
		if count.Next.Val == val {
			count.Next = count.Next.Next
		} else {
			count = count.Next
		}
	}
	return result.Next
}
