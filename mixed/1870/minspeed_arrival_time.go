package leet1870

import (
	"math"
)

func minSpeedOnTime(dist []int, hour float64) int {
	totalD := 0
	for _, val := range dist {
		totalD += val
	}
	s := int(float64(totalD) / hour)
	e := 10000000
	ans := -1

	for s <= e {
		mid := (s + e) / 2
		//fmt.Printf("l1: s=%v mid=%v e=%v\n", s, mid, e)
		if canTravelWithSpeed(dist, hour, mid) {
			ans = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
		//fmt.Printf("l2: s=%v mid=%v e=%v\n", s, mid, e)
	}

	return ans

}

func canTravelWithSpeed(dist []int, hour float64, speedL int) bool {
	timeT := 0.0
	for i := 0; i < (len(dist) - 1); i++ {
		timeT += math.Ceil(float64(dist[i]) / float64(speedL))
		if timeT > hour {
			return false
		}
	}
	timeT += float64(dist[len(dist)-1]) / float64(speedL)
	if timeT > hour {
		return false
	}
	return true

}
