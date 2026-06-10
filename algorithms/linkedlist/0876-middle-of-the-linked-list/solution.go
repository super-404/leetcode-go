package solution

import . "leetcode_go/datastructures"

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
