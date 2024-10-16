package l1143

func longestCommonSubsequence(text1 string, text2 string) int {
	lcsf := 0
	if len(text1) <= len(text2) {
		memsl := make([][]int, len(text1))
		initMemArray(memsl, len(text2))
		lcsf = longestCS(text1, text2, memsl, len(text1)-1, len(text2)-1)
	} else {
		memsl := make([][]int, len(text2))
		initMemArray(memsl, len(text1))
		lcsf = longestCS(text2, text1, memsl, len(text2)-1, len(text1)-1)
	}
	return lcsf
}

func initMemArray(memarr [][]int, eachrowl int) {
	for rown := range memarr {
		crow := make([]int, eachrowl)
		for coln := range crow {
			crow[coln] = -1
		}
		memarr[rown] = crow
	}
}

func longestCS(stxt, ltxt string, memsl [][]int, inds, indl int) int {
	// base case
	if inds < 0 || indl < 0 {
		return 0
	}
	if memsl[inds][indl] != -1 {
		return memsl[inds][indl]
	}
	// main logic
	lcsint := 0
	lcsc := 0
	if stxt[inds] == ltxt[indl] {
		lcsint = 1
		lcsc = longestCS(stxt, ltxt, memsl, inds-1, indl-1)
	} else {
		// indl not selected
		lcsstxtp := longestCS(stxt, ltxt, memsl, inds, indl-1)
		// indl selected
		lcsltxtp := longestCS(stxt, ltxt, memsl, inds-1, indl)
		lcsc = max2num(lcsstxtp, lcsltxtp)
	}
	memsl[inds][indl] = lcsint + lcsc
	return memsl[inds][indl]
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}
