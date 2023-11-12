package leet2226

func maximumCandies(candies []int, k int64) int {
	s := 1
	total := 0
	for _, val := range candies {
		total += val
	}
	e := (total / int(k))
	maxCan := 0
	for s <= e {
		mid := (s + e) / 2
		if canAllChildrenGetCandiesWithNum(candies, mid, k) {
			maxCan = mid
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return maxCan

}

func canAllChildrenGetCandiesWithNum(candies []int, target int, k int64) bool {
	kidWithCandy := 0
	for _, val := range candies {
		kidWithCandy += (val / target)
		if kidWithCandy >= int(k) {
			return true
		}
	}
	return false
}
