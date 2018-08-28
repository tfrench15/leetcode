package main

import "fmt"

func main() {
	nodeA := new(ListNode)
	nodeA = &ListNode{value: 1, next: nil}
	nodeB := &ListNode{value: 8, next: nil}
	nodeA.next = nodeB
	for node := nodeA; node != nil; node = node.next {
		fmt.Println("NodeA is equal to ", node)
	}
	node1 := new(ListNode)
	node1 = &ListNode{value: 0, next: nil}
	retNode := addTwoNumbers(nodeA, node1)
	for node := retNode; node != nil; node = node.next {
		fmt.Println("retNode is equal to ", node)
	}
}

// func twoSum returns an array of indices of elements
// summing to target.
func twoSum(nums []int, target int) []int {
	var ret []int
	idx := 0
	for idx < len(nums)-1 {
		for i, num := range nums {
			if idx == i {
				continue
			}
			if nums[idx]+num == target {
				ret = append(ret, idx, i)
				return ret
			}
		}
		idx++
	}
	return nil
}

// func addTwoNumbers adds two numbers, the digits of which
// are represented as nodes in a singly-linked list.
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1.value == 0 && l1.next == nil {
		return l2
	}
	if l2.value == 0 && l2.next == nil {
		return l1
	}

	ln := new(ListNode)
	carry := 0
	cur := ln

	cur.value = (l1.value + l2.value + carry) % 10
	carry = (l1.value + l2.value) / 10
	l1, l2 = l1.next, l2.next

	for {
		// Case I: l1, l2 both != nil
		if l1 != nil && l2 != nil {
			cur.next = new(ListNode)
			cur = cur.next
			cur.value = (l1.value + l2.value + carry) % 10
			carry = (l1.value + l2.value + carry) / 10
			l1, l2 = l1.next, l2.next
			continue
		}

		// Case II: l1 == nil, l2 != nil
		if l1 == nil && l2 != nil {
			cur.next = new(ListNode)
			cur = cur.next
			cur.value = (l2.value + carry) % 10
			carry = (l2.value + carry) / 10
			l2 = l2.next
			continue
		}

		// Case III: l1 != nil, l2 == nil
		if l1 != nil && l2 == nil {
			cur.next = new(ListNode)
			cur = cur.next
			cur.value = (l1.value + carry) % 10
			carry = (l1.value + carry) / 10
			l1 = l1.next
			continue
		}

		// Case IV: l1, l2 both == nil
		if l1 == nil && l2 == nil {
			if carry != 0 {
				cur.next = new(ListNode)
				cur = cur.next
				cur.value = carry
				return ln
			}
			return ln
		}
	}
}

// ListNode implements a singly-linked list
type ListNode struct {
	value int
	next  *ListNode
}
