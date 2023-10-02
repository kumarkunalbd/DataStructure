package l2187

func MinimumTime(time []int, totalTrips int) int64 {
	maxT := 0
	totalInUnitT := 0.0
	for _, val := range time {
		if val > maxT {
			maxT = val
		}
		totalInUnitT += (1.0 / float64(val))
	}

	e := (totalTrips * maxT)
	s := int(float64(totalTrips) / float64(totalInUnitT))
	minT := -1
	for s <= e {
		mid := (s + e) / 2
		if isTripsPossibleInTime(time, totalTrips, mid) {
			minT = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	return int64(minT)
}

func isTripsPossibleInTime(time []int, totalTrips int, targetT int) bool {
	curTrips := 0
	for _, val := range time {
		curTrips += (targetT / val)
		if curTrips >= totalTrips {
			return true
		}
	}

	return false
}
