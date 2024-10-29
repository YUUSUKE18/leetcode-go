package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	// リストが空か要素が1つしかない場合
	if head == nil || head.Next == nil {
		return false
	}

	// 2つのポインタを用意
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {

		if slow == fast {
			return true
		}

		slow = slow.Next

		fast = fast.Next.Next
	}

	// fastがnilに到達した場合は循環が存在しない
	return false
}
