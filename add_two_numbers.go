package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l := &ListNode{}
	if l1.Val+l2.Val < 10 {
		l.Val = l1.Val + l2.Val
		l.Next = addTwoNumbers(l1.Next, l2.Next)
	} else {
		l.Val = l1.Val + l2.Val - 10
		tmp := &ListNode{Val: 1}
		l.Next = addTwoNumbers(l1.Next, addTwoNumbers(tmp, l2.Next))
	}

	return l
}

//func main() {
//	l1 := &ListNode{Val: 2}
//	l1.Next = &ListNode{Val: 4}
//	l1.Next.Next = &ListNode{Val: 3}
//	l2 := &ListNode{Val: 5}
//	l2.Next = &ListNode{Val: 6}
//	l2.Next.Next = &ListNode{Val: 4}
//
//	l := addTwoNumbers(l1, l2)
//
//	var first = true
//	for ; l != nil; l = l.Next {
//		if !first {
//			fmt.Print(" -> ")
//		}
//		fmt.Print(l.Val)
//		first = false
//	}
//}
