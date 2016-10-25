package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var result, tmp *ListNode
	for {
		var v int
		if l1 != nil && l2 != nil {
			v = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			v = l1.Val
			l1 = l1.Next
		} else if l2 != nil {
			v = l2.Val
			l2 = l2.Next
		} else {
			if carry > 0 {
				tmp.Next = &ListNode{carry, nil}
			}
			break
		}

		v = v + carry
		carry = 0
		if v >= 10 {
			carry = v / 10
			v = v % 10
		}
		if result == nil {
			result = &ListNode{v, nil}
			tmp = result
			continue
		}
		tmp.Next = &ListNode{v, nil}
		tmp = tmp.Next
	}
	return result
}
