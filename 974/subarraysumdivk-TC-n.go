func subarraysDivByK(nums []int, k int) int {
	freqMapPsSum := make(map[int]int)
	freqMapPsSum[0] = 1
	curSum := 0
	count := 0
	for _, val := range nums {
		curSum += val
		modVal := curSum % k
		if modVal < 0 {
			modVal += k
		}
		count += freqMapPsSum[modVal]
		freqMapPsSum[modVal] = freqMapPsSum[modVal] + 1
	}
	return count
}