package hashing

func canPermutePalindrome(s string) bool {
	mapPresence := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		if mapPresence[s[i]] {
			delete(mapPresence, s[i])
		} else {
			mapPresence[s[i]] = true
		}
	}

	if len(mapPresence) <= 1 {
		return true
	}

	return false
}
