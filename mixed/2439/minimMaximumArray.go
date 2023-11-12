func minimizeArrayValue(nums []int) int {
	curSum := 0
	minMax := math.MinInt
	for i, val := range nums {
		curSum += val
		curAvg := int(math.Ceil(float64(curSum) / float64(i+1)))
		//fmt.Printf("i=%v val=%v curAvg=%v minMax=%v\n",i,val,curAvg,minMax)
		if curAvg > minMax {
			minMax = curAvg
		}
	}

	return minMax
}