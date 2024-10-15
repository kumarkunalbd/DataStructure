package l46

func permute(nums []int) [][]int {
	anAns := make([][]int, 0)
	inclusionI := make(map[int]bool, 0)
	curArr := make([]int, 0)
	anAns = permutationsWithCurInput(nums, inclusionI, curArr, -1, anAns)
	return anAns
}

func permutationsWithCurInput(nums []int, includeI map[int]bool, curArr []int, curI int, ans [][]int) [][]int {

	// base case
	if len(curArr) == len(nums) {
		dst := make([]int, len(curArr))
		copy(dst, curArr)
		ans = append(ans, dst)
		//fmt.Printf("After Addition =%v curArr=%v\n", ans,curArr)
		return ans
	}

	// cmd logic
	for i := 0; i < len(nums); i++ {
		if i != curI {
			if _, ok := includeI[i]; !ok {
				curArr = append(curArr, nums[i])
				includeI[i] = true
				ans = permutationsWithCurInput(nums, includeI, curArr, i, ans)
				curArr = curArr[:len(curArr)-1]
				delete(includeI, i)
			}
		}
	}
	return ans
}
