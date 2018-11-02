package answers

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	ln := new(ListNode)
	carry := 0
	cur := ln

	cur.Val = (l1.Val + l2.Val + carry) % 10
	carry = (l1.Val + l2.Val) / 10
	l1, l2 = l1.Next, l2.Next

	for {
		// Case I: l1, l2 both != nil
		if l1 != nil && l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			cur.Val = (l1.Val + l2.Val + carry) % 10
			carry = (l1.Val + l2.Val + carry) / 10
			l1, l2 = l1.Next, l2.Next
			continue
		}

		// Case II: l1 == nil, l2 != nil
		if l1 == nil && l2 != nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			cur.Val = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2 = l2.Next
			continue
		}

		// Case III: l1 != nil, l2 == nil
		if l1 != nil && l2 == nil {
			cur.Next = new(ListNode)
			cur = cur.Next
			cur.Val = (l1.Val + carry) % 10
			carry = (l1.Val + carry) / 10
			l1 = l1.Next
			continue
		}

		// Case IV: l1, l2 both == nil
		if l1 == nil && l2 == nil {
			if carry != 0 {
				cur.Next = new(ListNode)
				cur = cur.Next
				cur.Val = carry
				return ln
			}
			return ln
		}
	}
}
