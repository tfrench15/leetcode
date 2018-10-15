package answers

import "testing"

func TestMergeTwoLists(t *testing.T) {
	var (
		lA = new(ListNode)
		lB = new(ListNode)
		lC = new(ListNode)
		lX = new(ListNode)
		lY = new(ListNode)
		lZ = new(ListNode)
	)
	lA.Val, lA.Next = 1, lB
	lB.Val, lB.Next = 2, lC
	lC.Val, lC.Next = 4, nil
	lX.Val, lX.Next = 1, lY
	lY.Val, lY.Next = 3, lZ
	lZ.Val, lZ.Next = 4, nil
	mergedList := mergeTwoLists(lA, lX)
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	if mergedList != expected {
		t.Errorf("Error: expected %v, got %v", expected, mergedList)
	}
}
