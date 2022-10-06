package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	outputList := new(ListNode)
	for node := outputList; l1 != nil || l2 != nil || sum > 0; node = node.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{sum % 10, nil}
		sum /= 10
	}
	return outputList.Next
}
