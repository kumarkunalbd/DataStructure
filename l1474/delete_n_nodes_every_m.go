package l1474

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNodes(head *ListNode, m int, n int) *ListNode {
	var prev *ListNode
	ptr := head

	for ptr != nil {
		for i := 0; i < m && ptr != nil; i++ {
			prev = ptr
			ptr = (*ptr).Next
		}
		for i := 0; i < n && ptr != nil; i++ {
			ptr = (*ptr).Next
		}

		(*prev).Next = ptr
	}

	return head
}

func NewLinkedList(input []int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for i, val := range input {
		if i == 0 {
			head = &ListNode{val, nil}
			prev = head
		} else {
			ptr := &ListNode{val, nil}
			(*prev).Next = ptr
			prev = ptr
		}
	}
	return head
}
