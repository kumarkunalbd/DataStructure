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

type PairDiffNum struct {
	Diff float64
	Val  int
}

type MaxHeapTree struct {
	maxSize int
	Values  []PairDiffNum
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	ansArr := make([]int, 0)
	ansArr = closetKValuesWithSlice(root, target, k, ansArr)

	return ansArr

}

func closetKValuesWithSlice(root *TreeNode, target float64, k int, ans []int) []int {
	// base case
	if root == nil {
		return ans
	}
	// main logic
	ans = closetKValuesWithSlice(root.Left, target, k, ans)
	if len(ans) < k {
		ans = append(ans, root.Val)
	} else {
		curDiff := math.Abs(float64(root.Val) - target)
		diffFirst := math.Abs(float64(ans[0]) - target)
		if curDiff < diffFirst {
			ans = append(ans, root.Val)
			ans = ans[1:]
		} else {
			return ans
		}
	}
	ans = closetKValuesWithSlice(root.Right, target, k, ans)
	return ans
}

func closestKPairDiffNum(root *TreeNode, target float64, k int, maxHeap *MaxHeapTree) {
	// base case
	if root == nil {
		return
	}
	// main logic

	curDiff := math.Abs(float64(root.Val) - target)
	if len(maxHeap.Values) < k {
		//fmt.Printf("AddAlways root=%v heap=%v\n",root.Val,maxHeap.Values)
		maxHeap.addValueHeapify(curDiff, root.Val)
		closestKPairDiffNum(root.Left, target, k, maxHeap)
		closestKPairDiffNum(root.Right, target, k, maxHeap)
	} else {
		//fmt.Printf("CheckAdd root=%v heap=%v\n",root.Val,maxHeap.Values)
		topVal := maxHeap.peek()
		if curDiff < topVal.Diff {
			maxHeap.removeValueHeapify()
			//fmt.Printf("removal: heap=%v\n",maxHeap.Values)
			maxHeap.addValueHeapify(curDiff, root.Val)
			//fmt.Printf("addition: heap=%v\n",maxHeap.Values)
		}
		//fmt.Printf("root=%v\n",root.Val)
		if float64(root.Val) > target {
			closestKPairDiffNum(root.Left, target, k, maxHeap)
		} else {
			closestKPairDiffNum(root.Right, target, k, maxHeap)
		}
	}
	//fmt.Printf("final root=%v heap=%v\n",root.Val,maxHeap.Values)
}

func NewMaxHeap(size int) *MaxHeapTree {
	return &MaxHeapTree{
		maxSize: size,
		Values:  make([]PairDiffNum, 0),
	}
}

func (maxHeap *MaxHeapTree) addValueHeapify(diff float64, val int) {
	maxHeap.Values = append(maxHeap.Values, PairDiffNum{diff, val})
	curI := len(maxHeap.Values) - 1
	for curI > 0 {
		parentI := (curI - 1) / 2
		if maxHeap.Values[parentI].Diff < maxHeap.Values[curI].Diff {
			maxHeap.Values[parentI], maxHeap.Values[curI] = maxHeap.Values[curI], maxHeap.Values[parentI]
			curI = parentI
		} else {
			break
		}
	}
}

func (maxHeap *MaxHeapTree) removeValueHeapify() bool {
	if !maxHeap.isEmpty() {
		maxHeap.Values[0], maxHeap.Values[len(maxHeap.Values)-1] = maxHeap.Values[len(maxHeap.Values)-1], maxHeap.Values[0]
		maxHeap.Values = maxHeap.Values[0:(len(maxHeap.Values) - 1)]
		curI := 0
		for curI < len(maxHeap.Values) {
			maxI := curI
			leftC := (2 * curI) + 1
			if leftC < len(maxHeap.Values) && maxHeap.Values[leftC].Diff > maxHeap.Values[maxI].Diff {
				maxI = leftC
			}
			rightC := (2 * curI) + 2
			if rightC < len(maxHeap.Values) && maxHeap.Values[rightC].Diff > maxHeap.Values[maxI].Diff {
				maxI = rightC
			}
			if maxI == curI {
				break
			} else {
				maxHeap.Values[curI], maxHeap.Values[maxI] = maxHeap.Values[maxI], maxHeap.Values[curI]
				curI = maxI
			}
		}
		return true
	}
	return false
}

func (maxHeap *MaxHeapTree) peek() *PairDiffNum {
	/*if len(maxHeap.Values)>0{
	      pair := maxHeap.Values[0]
	      return &pair,nil
	  }
	  return nil,fmt.Errof("heap is empty")*/
	pair := maxHeap.Values[0]
	return &pair
}

func (maxHeap MaxHeapTree) isEmpty() bool {
	if len(maxHeap.Values) == 0 {
		return true
	}
	return false
}
