package l78

import "slices"

func Subsets(nums []int) [][]int {
	curSet := make([]int, 0)
	var ans [][]int
	ans = dfsSubset(nums, len(nums)-1, curSet, ans)
	return ans
}

func dfsSubset(nums []int, curI int, curArr []int, ansArr [][]int) [][]int {
	// base case
	if curI == -1 {
		//fmt.Printf("curI=%v curArr=%v ans=%v\n", curI,curArr,ansArr)
		cloneArr := make([]int, len(curArr))
		_ = copy(cloneArr, curArr)
		slices.Sort(cloneArr)
		ansArr = append(ansArr, cloneArr)
		return ansArr
	}
	// cmd logic
	ansArr = dfsSubset(nums, curI-1, curArr, ansArr)
	curArr = append(curArr, nums[curI])
	ansArr = dfsSubset(nums, curI-1, curArr, ansArr)
	return ansArr
}
