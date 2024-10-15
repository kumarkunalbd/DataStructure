package validpath

/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer
 * @input D : Integer
 * @input E : Integer array
 * @input F : Integer array
 *
 * @Output string.
 */

import (
	//"fmt"
	"math"
)

func vpath(A int, B int, D int, E []int, F []int) string {

	memPathMat := prepPathMat(A, B, D, E, F)
	//fmt.Printf("memPath=%v\n",memPathMat)
	if memPathMat[0][0] == 0 || memPathMat[A][B] == 0 {
		return "NO"
	}
	rowM := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	colM := []int{0, 1, 1, 1, 0, -1, -1, -1}
	/*vexist := recValidPath(D,memPathMat,A,B,rowM,colM)
	  if vexist {
	      return "YES"
	  }
	  return "NO"*/
	aq := newqueue()
	aq.enq(&cord{0, 0})
	memPathMat[0][0] = 1

	for !aq.isempty() {
		outcord := aq.dq()
		for i := 0; i < len(rowM); i++ {
			nr := outcord.x + rowM[i]
			nc := outcord.y + colM[i]
			if isrowvalid(nr, memPathMat) && iscolvalid(nc, memPathMat) && memPathMat[nr][nc] == -1 {
				if nr == A && nc == B {
					return "YES"
				}
				memPathMat[nr][nc] = 1
				aq.enq(&cord{
					x: nr,
					y: nc,
				})
			}
		}
	}

	return "NO"
}

type cord struct {
	x int
	y int
}

type queue []*cord

func newqueue() *queue {
	arr := make(queue, 0)
	return &arr
}

func (q *queue) enq(cord *cord) {
	*q = append(*q, cord)
}

func (q *queue) dq() *cord {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *queue) isempty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func prepPathMat(x, y, R int, xCirC, yCirC []int) [][]int {
	memPathR := make([][]int, x+1)
	for i, row := range memPathR {
		row = make([]int, y+1)
		memPathR[i] = row
		for j := range row {
			//fmt.Printf("i=%v row=%v j=%v\n",i,row,j)
			if isCordUnderCircle(i, j, R, xCirC, yCirC) {
				memPathR[i][j] = 0
			} else {
				memPathR[i][j] = -1
			}
		}
	}
	return memPathR
}

func isCordUnderCircle(curx, cury, R int, xCirC, yCirC []int) bool {
	//fmt.Printf("isCordUnderCircle: curx=%v cury=%v R=%v\n",curx,cury,R)
	for i, curCx := range xCirC {
		curCy := yCirC[i]
		sqDist := math.Sqrt(float64((curx-curCx)*(curx-curCx)) + float64((cury-curCy)*(cury-curCy)))
		//fmt.Printf("curCx=%v curCy=%v sqDist=%v\n",curCx,curCy,sqDist)
		if sqDist <= float64(R) {
			return true
		}
	}
	return false
}

func recValidPath(R int, dp [][]int, curx, cury int, rowM, colM []int) bool {
	//fmt.Printf("recValidPath: curx=%v cury=%v\n",curx,cury)
	// base case
	if dp[curx][cury] != -1 {
		if dp[curx][cury] == 1 {
			return true
		}
		return false
	}
	// cmd logic
	for i := 0; i < len(rowM); i++ {
		nx := curx + rowM[i]
		ny := cury + colM[i]
		if isrowvalid(nx, dp) && iscolvalid(ny, dp) {
			vpath := recValidPath(R, dp, nx, ny, rowM, colM)
			if vpath {
				dp[curx][cury] = 1
				return true
			}
		}
	}
	dp[curx][cury] = 0
	return false
}

func isrowvalid(row int, memMat [][]int) bool {
	if row >= 0 && row < len(memMat) {
		return true
	}
	return false
}

func iscolvalid(col int, memMat [][]int) bool {
	if col >= 0 && col < len(memMat[0]) {
		return true
	}
	return false
}
