func grayCode(n int) []int {
	ans := []int{}
	ans = grayCodeHelper(n-1, ans)
	fmt.Printf("ans:%#v", ans)
	return ans
}

func grayCodeHelper(bitPos int, res []int) []int {
	// base case
	if bitPos == 0 {
		res = append(res, []int{0, 1}...)
		return res
	}

	// main logic
	res = grayCodeHelper(bitPos-1, res)
	sLength := len(res)
	for i := sLength - 1; i >= 0; i-- {
		res = append(res, res[i]+(1<<bitPos))
	}
	return res
}