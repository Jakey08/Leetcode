/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func deleteDuplicates(head *ListNode) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	cur := head
	for cur != nil {
		next := cur.Next
		for next != nil && next.Val == cur.Val {
			next = next.Next
		}
		cur.Next = next
		cur = next
	}
	return prev.Next