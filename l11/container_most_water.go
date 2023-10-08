package l11

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {
		var CA int
		if height[l] < height[r] {
			CA = (height[l]) * (r - l)
			l++
		} else {
			CA = (height[r]) * (r - l)
			r--
		}

		if CA > maxArea {
			maxArea = CA
		}
	}
	return maxArea
}
