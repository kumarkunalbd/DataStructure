/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package l272

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Dequeue struct {
	Values []int
	length int
	ind    int
}

func NewDequeWithLen(size int) *Dequeue {
	return &Dequeue{
		Values: make([]int, 0),
		length: size,
		ind:    -1,
	}
}

func (aDQ *Dequeue) AddLast(val int) {
	aDQ.Values = append(aDQ.Values, val)
	aDQ.ind++
}

func (aDQ *Dequeue) RemoveFront() int {
	ele := aDQ.Values[0]
	aDQ.Values = aDQ.Values[1:]
	aDQ.ind--
	return ele
}

func (aDQ *Dequeue) IsFull() bool {
	if len(aDQ.Values) == aDQ.length {
		return true
	}
	return false
}

func (aDQ *Dequeue) PeekFront() int {
	return aDQ.Values[0]
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	ansDQ := NewDequeWithLen(k)
	closetKValuesWithSlice(root, target, k, ansDQ)

	return ansDQ.Values
}

func closetKValuesWithSlice(root *TreeNode, target float64, k int, adq *Dequeue) {
	// base case
	if root == nil {
		return
	}
	// cmd logic
	closetKValuesWithSlice(root.Left, target, k, adq)
	curDiff := math.Abs(float64(root.Val) - target)
	if adq.IsFull() {
		firstdiff := math.Abs(float64(adq.PeekFront()) - target)
		if curDiff >= firstdiff {
			return
		} else {
			_ = adq.RemoveFront()
			adq.AddLast(root.Val)
		}
	} else {
		adq.AddLast(root.Val)
	}
	closetKValuesWithSlice(root.Right, target, k, adq)
}
