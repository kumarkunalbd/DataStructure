package l516

func longestPalindromeSubseq(s string) int {
	memLP := make([][]int, len(s))
	for rown := range memLP {
		crow := make([]int, len(s))
		//crow[rown] = 1
		memLP[rown] = crow
	}
	//fmt.Printf("memLP=%v\n",memLP)
	return longestPS(s, 0, len(s)-1, memLP)
}

func longestPS(s string, indl, indr int, memLP [][]int) int {
	//fmt.Printf("memLP=%v indl=%v indr=%v\n",memLP,indl,indr)
	if indl > indr {
		return 0
	}
	if indl == indr {
		return 1
	}

	if memLP[indl][indr] != 0 {
		return memLP[indl][indr]
	}

	// main logic
	lpsinit := 0
	lpsc := 0
	if s[indl] == s[indr] {
		lpsinit = 2
		lpsc = longestPS(s, indl+1, indr-1, memLP)
	} else {
		// if r is not selected
		lpscns := longestPS(s, indl, indr-1, memLP)
		// if r is selected
		lpscs := longestPS(s, indl+1, indr, memLP)
		lpsc = max2num(lpscns, lpscs)
	}
	memLP[indl][indr] = lpsinit + lpsc
	return memLP[indl][indr]
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}
