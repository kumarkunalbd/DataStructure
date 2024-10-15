package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	return mergeSortListSpaceOpt(head)
}

func mergeSortList(mapNodes map[int]*ListNode, s, e int) *ListNode {
	// base case
	if s == e {
		return mapNodes[s]
	}

	// cmd logic
	mid := s + (e-s+1)/2
	aNode := mapNodes[mid-1]
	aNode.Next = nil
	lHead := mergeSortList(mapNodes, s, mid-1)
	rHead := mergeSortList(mapNodes, mid, e)
	head := mergeList(lHead, rHead)
	return head
}

func mergeSortListSpaceOpt(startNode *ListNode) *ListNode {
	// base case
	if startNode.Next == nil {
		return startNode
	}

	// cmd logic
	s, f := startNode, startNode
	prev := &ListNode{
		Next: nil,
	}
	for f != nil && f.Next != nil {
		prev = s
		s = s.Next
		f = f.Next.Next
	}
	prev.Next = nil
	lHead := mergeSortListSpaceOpt(startNode)
	rHead := mergeSortListSpaceOpt(s)

	return mergeList(lHead, rHead)

}

func mergeList(lHead, rHead *ListNode) *ListNode {
	if lHead == nil {
		return rHead
	}
	if rHead == nil {
		return lHead
	}
	dummy := &ListNode{
		Next: nil,
	}
	prev := dummy
	for lHead != nil && rHead != nil {
		if lHead.Val <= rHead.Val {
			prev.Next = lHead
			lHead = lHead.Next

		} else {
			prev.Next = rHead
			rHead = rHead.Next
		}
		prev = prev.Next
	}

	if lHead != nil {
		prev.Next = lHead
	}
	if rHead != nil {
		prev.Next = rHead
	}

	return dummy.Next
}
