package solution

import . "leetcode_go/data_structures"

// LeetCode 86. Partition List
// 将链表分割，使得所有小于 x 的节点都在大于或等于 x 的节点之前
func partition(head *ListNode, x int) *ListNode {
	lessDummy := &ListNode{}
	greaterDummy := &ListNode{}
	less := lessDummy
	greater := greaterDummy

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}

	greater.Next = nil
	less.Next = greaterDummy.Next
	return lessDummy.Next
}
