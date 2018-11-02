package answers

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	// Easy case, both lists are nil
	if l1 == nil && l2 == nil {
		return nil
	}

	var (
		ret = new(ListNode)
		cur = new(ListNode)
	)
	ret = cur

	for {
		if l1 == nil && l2 == nil {
			return ret
		}
		if l1 == nil && l2 != nil {
			cur.Val = l2.Val
			l2 = l2.Next
			if l2 == nil {
				return ret
			}
			cur.Next = new(ListNode)
			cur = cur.Next
			continue
		}
		if l1 != nil && l2 == nil {
			cur.Val = l1.Val
			l1 = l1.Next
			if l1 == nil {
				return ret
			}
			cur.Next = new(ListNode)
			cur = cur.Next
			continue
		}
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				cur.Val = l1.Val
				cur.Next = new(ListNode)
				cur, l1 = cur.Next, l1.Next
			} else {
				cur.Val = l2.Val
				cur.Next = new(ListNode)
				cur, l2 = cur.Next, l2.Next
			}
		}
	}
}
