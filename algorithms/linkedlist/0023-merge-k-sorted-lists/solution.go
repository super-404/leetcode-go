package solution

import . "leetcode_go/datastructures"

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	p := dummy
	count := 0
	for count < len(lists)-1 {
		minH := lists[0]
		i := 0

		for k, h := range lists {
			if h == nil {
				count++
				continue
			}
			if minH == nil || h.Val < minH.Val {
				minH = h
				i = k
			}
		}
		if minH == nil {
			break
		}
		p.Next = minH
		p = minH
		lists[i] = minH.Next
		count = 0
	}
	for _, h := range lists {
		if h != nil {
			p.Next = h
		}
	}

	return dummy.Next
}
