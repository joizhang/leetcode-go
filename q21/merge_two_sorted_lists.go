package q21

type ListNode struct {
	Val int
	Next *ListNode
}

// 21. Merge Two Sorted Lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := new(ListNode)
	if l1.Val <= l2.Val {
		dummy.Next = l1
		l1 = l1.Next
	} else {
		dummy.Next = l2
		l2 = l2.Next
	}
	p := dummy.Next
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			p.Next = l2
			break
		}
		if l1 != nil && l2 == nil {
			p.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	return dummy.Next
}
