func candy(ratings []int) int {
	l := 0
	r := 0
	minT := 1
	prevC := 1
	upperCap := 0

	for r = 1; r < len(ratings); r++ {
		if ratings[r] < ratings[r-1] {
			if prevC == 1 {
				if (r - l) < upperCap {
					minT += (r - l - 1) + 1
				} else {
					minT += (r - l) + 1
				}
			} else {
				minT += 1
			}
			prevC = 1
		} else if ratings[r] == ratings[r-1] {
			minT += 1
			l = r
			prevC = 1
			upperCap = prevC
		} else {
			minT += prevC + 1
			prevC = prevC + 1
			l = r
			upperCap = prevC
		}
	}
	return minT
}