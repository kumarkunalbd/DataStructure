func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0
	minL := math.MaxFloat64
	curSum := 0
	for r < len(nums) {
		curSum += nums[r]
		if curSum < target {
			r++
			continue
		}
		for curSum >= target {
			curSum -= nums[l]
			l++
		}
		minL = math.Min(float64(r)-float64(l)+2.0, minL)
		//fmt.Printf("curSum=%v r=%v l=%v minL=%v\n",curSum,r,l,minL)
		r++
	}
	if minL == math.MaxFloat64 {
		return 0
	} else {
		return int(minL)
	}

}