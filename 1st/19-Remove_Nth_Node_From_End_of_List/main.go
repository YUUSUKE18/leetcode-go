package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{0, head}
	top := result
	bottom := result

	for i := 1; i <= n; i++ {
		if top.Next == nil {
			return head
		}
		top = top.Next
	}

	for top.Next != nil {
		top = top.Next
		bottom = bottom.Next
	}

	if bottom.Next != nil {
		bottom.Next = bottom.Next.Next
	}

	return result.Next
}
