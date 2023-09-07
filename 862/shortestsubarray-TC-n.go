func shortestSubarray(nums []int, k int) int {
	psSum := make([]int, len(nums)+1)
	for i, val := range nums {
		psSum[i+1] = psSum[i] + val
	}
	//fmt.Printf("psSum=%v\n",psSum)

	minLength := float64(len(nums) + 1)
	aQueue := make([]int, 0)
	for r, val := range psSum {
		for len(aQueue) != 0 && val < psSum[aQueue[len(aQueue)-1]] {
			aQueue = aQueue[:len(aQueue)-1]
		}

		for len(aQueue) != 0 && val >= (psSum[aQueue[0]]+k) {
			minLength = math.Min(float64(r-aQueue[0]), minLength)
			aQueue = aQueue[1:]
		}
		aQueue = append(aQueue, r)
	}

	if minLength == float64(len(nums)+1) {
		return -1
	}
	return int(minLength)

}