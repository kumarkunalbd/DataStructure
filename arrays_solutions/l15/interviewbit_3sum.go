package l15

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer
 */

import (
	"sort"
	//"fmt"
	"math"
)

type SortedInt []int

func (arr SortedInt) Len() int {
	return len(arr)
}
func (arr SortedInt) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr SortedInt) Less(i, j int) bool {
	return arr[i] <= arr[j]
}
func threeSumClosest(A []int, B int) int {
	sortedArr := SortedInt(A)
	sort.Sort(sortedArr)
	closest := math.MaxInt
	ans := math.MinInt
	for i, val := range A {
		if i > 0 && val == A[i-1] {
			continue
		}
		l := i + 1
		r := len(A) - 1
		for l < r {
			curDiff := (A[l] + A[r] + val) - B
			if getAbs(curDiff) < closest {
				closest = getAbs(curDiff)
				ans = A[l] + A[r] + val
				if closest == 0 {
					return A[l] + A[r] + val
				}
			}
			if curDiff < 0 {
				l++
			} else {
				r--
			}
		}
	}
	//fmt.Printf("Sorted Arr:%v\n",A)
	return ans
}

func getAbs(num int) int {
	if num < 0 {
		return 0 - num
	}
	return num
}
