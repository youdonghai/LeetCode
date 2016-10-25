package addTwoNumbers

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	r := AddTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}})
	if r == nil || r.Val != 7 ||
		r.Next == nil || r.Next.Val != 0 ||
		r.Next.Next == nil || r.Next.Next.Val != 8 {
		t.Error("addTwoNumbers error")
	}
}
