package _25

func maxSubArrayLen(nums []int, k int) int {
	mapIPos := make(map[int]int, 0)
	mapIPos[0] = -1
	maxLen := 0
	curSum := 0
	for i, val := range nums {
		curSum += val
		//fmt.Printf("i=%v curSum=%v mapPos=%v\n",i,curSum,mapIPos)
		if iPos, ok := mapIPos[curSum-k]; ok {
			if (i - iPos) > maxLen {
				maxLen = (i - iPos)
			}
		}
		if _, ok := mapIPos[curSum]; !ok {
			mapIPos[curSum] = i
		}
	}

	return maxLen
}
