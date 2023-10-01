func minimizeArrayValue(nums []int) int {
	s := math.MaxInt
	e := math.MinInt

	for _, val := range nums {
		if val < s {
			s = val
		}
		if val > e {
			e = val
		}
	}

	ans := e
	for s <= e {
		mid := (s + e) / 2
		fmt.Printf("l1: s=%v mid=%v e=%v\n", s, mid, e)
		if canMaximumValueBe(nums, mid) {
			ans = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
		fmt.Printf("l2: s=%v mid=%v e=%v ans=%v\n", s, mid, e, ans)
	}
	return ans
}

func canMaximumValueBe(nums []int, target int) bool {
	if nums[0] > target {
		return false
	}
	prevMinPossible := nums[0]
	for i := 1; i < len(nums); i++ {
		curMinPoss := nums[i] - (target - prevMinPossible)
		if curMinPoss > target {
			return false
		}
		prevMinPossible = curMinPoss
	}
	return true
}