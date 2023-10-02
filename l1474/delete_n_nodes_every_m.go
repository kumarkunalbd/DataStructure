package l1474

import "fmt"

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
	tmp := head
	//var lastNonDelete *ListNode
	lastNonDelete := head
	count := 1
	for tmp != nil {
		fmt.Printf("l1: count=%v tmp=%v last=%v\n", count, tmp, lastNonDelete)
		if count%(m+n) == 0 {
			(*lastNonDelete).Next = (*tmp).Next
		}
		if count%(m+n) < m {
			lastNonDelete = (*tmp).Next
			tmp = (*tmp).Next
		} else {
			tmp = (*tmp).Next
			(*lastNonDelete).Next = nil
		}
		count++
		fmt.Printf("l2: count=%v tmp=%v last=%v\n", count, tmp, lastNonDelete)
	}
	return head
}
