package solution

import . "leetcode_go/datastructures"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1 := headA
	p2 := headB
	len1 := 0
	len2 := 0
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil {
		len2++
		p2 = p2.Next
	}
	p1 = headA
	p2 = headB
	if len1 > len2 {
		i := 0
		for i < len1-len2 {
			p1 = p1.Next
			i++
		}
	} else {
		i := 0
		for i < len2-len1 {
			p2 = p2.Next
			i++
		}
	}
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil

}
