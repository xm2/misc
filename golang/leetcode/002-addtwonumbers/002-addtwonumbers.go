package addtwonumbers

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

/*ListNode Definition for singly-linked list.*/
type ListNode struct {
	Val  uint
	Next *ListNode
}

func makeNodeListFromNum(num uint) *ListNode {
	var ret *ListNode
	r := []uint{}
	for {
		if num < 10 {
			r = append(r, num)
			break
		}

		r = append(r, num%10)
		num = num / 10
	}

	for i := len(r) - 1; i >= 0; i-- {
		n := &ListNode{r[i], nil}
		n.Next = ret
		ret = n
	}

	return ret
}

func nodeListToNum(l *ListNode) uint {
	factor := uint(1)
	ret := uint(0)
	for {
		ret = ret + l.Val*factor

		factor *= 10

		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	ret := nodeListToNum(l1) + nodeListToNum(l2)

	r := makeNodeListFromNum(ret)

	return r
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{0, nil} //dummy node
	curr := r
	carry := uint(0)
	for {
		if l1 == nil && l2 == nil {
			break
		}
		v1, v2 := uint(0), uint(0)
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{1, nil}
	}

	return r.Next

}
