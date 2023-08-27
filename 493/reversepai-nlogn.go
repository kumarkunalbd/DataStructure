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

	mergeArr := make([]int, end-start)
	l = start
	r = mid
	counter := 0

	for l < mid && r < end {
		if nums[l] <= nums[r] {
			mergeArr[counter] = nums[l]
			l++
		} else {
			mergeArr[counter] = nums[r]
			r++
		}
		counter++
	}

	if l == mid {
		copy(nums[start:(start+counter)], mergeArr[0:counter])
	} else {
		copy(nums[(start+counter):end], nums[l:mid])
		copy(nums[start:(start+counter)], mergeArr[:counter])
	}

	return rCount

}