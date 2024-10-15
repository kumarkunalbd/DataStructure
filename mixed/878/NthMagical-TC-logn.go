func nthMagicalNumber(n int, a int, b int) int {
	m := 1000000007
	gcd := getGCD(a, b)
	lcm := (a / gcd) * b

	s := 0
	e := n * int(math.Min(float64(a), float64(b)))

	for s < e {
		mid := (s + e) / 2

		if (mid/a + mid/b - mid/lcm) < n {
			s = mid + 1
		} else {
			e = mid
		}
	}

	return int(s % m)

}

func getGCD(a, b int) int {
	// base vase
	if a == 0 {
		return b
	}
	// cmd logic
	if a < b {
		a, b = b, a
	}
	gcd := getGCD(a%b, b)
	return gcd
}