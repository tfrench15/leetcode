package main

import "fmt"

func main() {
	sub := lengthOfLongestSubstring("au")
	fmt.Println(sub)
	test := "hello"
	fmt.Println(test[0:4])
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

func lengthOfLongestSubstring(s string) int {
	// Easy cases
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var lengths []int
	var length int
	for i := 0; i < len(s)-1; i++ {
		seen := make(map[byte]bool)
		length = 0
		for j := i; j <= len(s)-1; j++ {
			_, ok := seen[s[j]]
			if !ok {
				length++
				seen[s[j]] = true
			} else {
				lengths = append(lengths, length)
				break
			}
		}
	}
	if length != 0 {
		lengths = append(lengths, length)
	}
	if len(lengths) == 0 {
		return 0
	}

	max := maxInt(lengths)
	return max
}

/*
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		if len(s) == 0 {
			return 0
		}
		return 1
	}
	lengths := []int{}
	for i := 0; i <= len(s)-1; i++ {
		seen := make(map[byte]bool)
		length := 0
		for j := i; j <= len(s)-1; j++ {
			_, ok := seen[s[j]]
			if !ok {
				length++
				fmt.Printf("The length is now %v\n", length)
				seen[s[j]] = true
				continue
			}
			lengths = append(lengths, length)
			break
		}
	}
	if len(lengths) == 0 {
		return 0
	}
	max := maxInt(lengths)
	return max
}
*/

func maxInt(slc []int) int {
	max := slc[0]
	for _, v := range slc {
		if v > max {
			max = v
		}
	}
	return max
}

// ListNode implements a singly-linked list
type ListNode struct {
	value int
	next  *ListNode
}
