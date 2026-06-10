package solution

import . "leetcode_go/datastructures"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	if head == nil {
		return head
	}
	p1 := dummy
	p2 := dummy
	for n != 0 {
		p1 = p1.Next
		n--
	}
	for p1 != nil {
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			break
		}
		p2 = p2.Next

	}
	if p2 != nil && p2.Next != nil {
		p2.Next = p2.Next.Next
	}
	return dummy.Next
}
