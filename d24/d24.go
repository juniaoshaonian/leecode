package d24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	ans := pre
	for head.Next != nil {
		next := head.Next.Next
		head.Next.Next = head
		pre.Next = head.Next
		head.Next = next
		pre = head
		head = head.Next
	}
	return ans.Next

}
