import (
	"fmt"
)

func reversePairs(nums []int) int {
	//fmt.Printf("original=%v\n",nums)
	ans := mergeSort(nums, 0, len(nums))
	//fmt.Printf("sorted=%#v\n", nums)

	return ans

}

func mergeSort(nums []int, start, end int) int {

	// base case
	if (end - start) <= 1 {
		return 0
	}

	// main logic
	mid := start + (end-start)/2
	leftC := mergeSort(nums, start, mid)
	rightC := mergeSort(nums, mid, end)
	curC := merge(nums, start, mid, end)

	return leftC + rightC + curC

}

func merge(nums []int, start, mid, end int) int {
	l := start
	r := mid
	rCount := 0
	for l < mid && r < end {
		if nums[l] > (2 * nums[r]) {
			rCount += (mid - l)
			r++
		} else {
			l++
		}
	}

	mergeArr := make([]int, 0)
	l = start
	r = mid
	for l < mid && r < end {
		if nums[l] <= nums[r] {
			mergeArr = append(mergeArr, nums[l])
			l++
		} else {
			mergeArr = append(mergeArr, nums[r])
			r++
		}
	}

	if l == mid {
		copy(nums[start:(start+len(mergeArr))], mergeArr[0:])
	} else {
		copy(nums[(start+len(mergeArr)):end], nums[l:mid])
		copy(nums[start:(start+len(mergeArr))], mergeArr[0:])
	}

	return rCount

}