package l51

import (
	"math"
	"strings"
)

type Point struct {
	row int
	col int
}

func SolveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	setVPoints := make(map[Point]bool, 0)
	curArr := make([]string, 0)
	ans = dfsNQueen(n, 0, curArr, ans, setVPoints)
	return ans
}

func dfsNQueen(n, curRow int, curArrString []string, ans [][]string, setVPoints map[Point]bool) [][]string {
	// base case
	if curRow == n {
		newArrString := make([]string, len(curArrString))
		copy(newArrString, curArrString)
		ans = append(ans, newArrString)
		//fmt.Printf("adding:  ans=%v\n", ans)
		return ans
	}
	// main logic
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteRune('.')
	}
	curPlace := sb.String()
	byteS := []rune(curPlace)
	//fmt.Printf("curRow=%v ans=%v setVPoints=%v curArrString=%v",curRow,ans,setVPoints,curArrString)
	for i := 0; i < n; i++ {
		isValidP := true
		for vPoint := range setVPoints {
			if i == vPoint.col {
				isValidP = false
				break
			}
			diffRow := int(math.Abs(float64(curRow - vPoint.row)))
			diffCol := int(math.Abs(float64(i - vPoint.col)))
			if diffRow == diffCol {
				isValidP = false
				break
			}
		}
		if isValidP {
			byteS[i] = 'Q'
			curArrString = append(curArrString, string(byteS))
			setVPoints[Point{curRow, i}] = true
			ans = dfsNQueen(n, curRow+1, curArrString, ans, setVPoints)
			byteS[i] = '.'
			curArrString = curArrString[:len(curArrString)-1]
			delete(setVPoints, Point{curRow, i})
		}
	}
	return ans
}
