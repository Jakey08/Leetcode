
func deleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cul.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return prev.Next
}