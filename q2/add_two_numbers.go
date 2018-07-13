package q2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2. Add Two Numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	temp := &ListNode{0, nil}
	res = temp
	carry := 0
	for l1 != nil || l2 != nil  {
		if l1 != nil && l2 != nil {
			temp.Next = &ListNode{(l1.Val + l2.Val + carry) % 10, nil}
			carry = (l1.Val + l2.Val + carry) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 != nil {
			temp.Next = &ListNode{(l2.Val + carry) % 10, nil}
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			temp.Next = &ListNode{(l1.Val + carry) % 10, nil}
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
		}
		temp = temp.Next
	}
	if carry != 0 {
		temp.Next = &ListNode{carry % 10, nil}
	}
	return res.Next
}
