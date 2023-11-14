package l502

import "math"

type IPO struct {
	Profit  int
	Capital int
}

type MaxHeap struct {
	Values  []IPO
	CurSize int
}

func (maxHeap *MaxHeap) AddElement(val IPO) {
	maxHeap.Values = append(maxHeap.Values, val)
	curI := len(maxHeap.Values) - 1
	for curI > 0 {
		parentI := (curI - 1) / 2
		if (maxHeap.Values[curI].Profit > maxHeap.Values[parentI].Profit) || (maxHeap.Values[curI].Profit == maxHeap.Values[parentI].Profit && maxHeap.Values[curI].Capital < maxHeap.Values[parentI].Capital) {
			maxHeap.Values[curI], maxHeap.Values[parentI] = maxHeap.Values[parentI], maxHeap.Values[curI]
			curI = parentI
		} else {
			break
		}
	}
	maxHeap.CurSize++
}

func (maxHeap *MaxHeap) RemoveElementTop() {
	maxHeap.RemoveElementAtIndex(0)
	/*maxHeap.Values[0],maxHeap.Values[maxHeap.CurSize-1] = maxHeap.Values[maxHeap.CurSize-1],maxHeap.Values[0]
	  maxHeap.Values = maxHeap.Values[:len(maxHeap.Values)-1]
	  maxHeap.CurSize--
	  curI := 0
	  for {
	      leftC := 2*curI+1
	      maxI := curI
	      if leftC < len(maxHeap.Values) && (maxHeap.Values[leftC].Profit>maxHeap.Values[curI].Profit || (maxHeap.Values[leftC].Profit == maxHeap.Values[curI].Profit && maxHeap.Values[leftC].Capital < maxHeap.Values[curI].Capital)){
	          maxI = leftC
	      }

	      rightC := 2*curI+2
	      if rightC < len(maxHeap.Values) && (maxHeap.Values[rightC].Profit>maxHeap.Values[maxI].Profit || (maxHeap.Values[rightC].Profit == maxHeap.Values[maxI].Profit && maxHeap.Values[rightC].Capital < maxHeap.Values[maxI].Capital)){
	          maxI = rightC
	      }
	      if maxI != curI{
	          maxHeap.Values[curI],maxHeap.Values[maxI] = maxHeap.Values[maxI],maxHeap.Values[curI]
	          curI = maxI
	      }else{
	          break
	      }
	  }*/
}

func (maxHeap *MaxHeap) RemoveElementAtIndex(ind int) {
	maxHeap.Values[ind], maxHeap.Values[maxHeap.CurSize-1] = maxHeap.Values[maxHeap.CurSize-1], maxHeap.Values[ind]
	maxHeap.Values = maxHeap.Values[:len(maxHeap.Values)-1]
	maxHeap.CurSize--
	curI := ind
	for curI < len(maxHeap.Values) {
		leftC := 2*curI + 1
		maxI := curI
		if leftC < len(maxHeap.Values) && (maxHeap.Values[leftC].Profit > maxHeap.Values[curI].Profit || (maxHeap.Values[leftC].Profit == maxHeap.Values[curI].Profit && maxHeap.Values[leftC].Capital < maxHeap.Values[curI].Capital)) {
			maxI = leftC
		}

		rightC := 2*curI + 2
		if rightC < len(maxHeap.Values) && (maxHeap.Values[rightC].Profit > maxHeap.Values[maxI].Profit || (maxHeap.Values[rightC].Profit == maxHeap.Values[maxI].Profit && maxHeap.Values[rightC].Capital < maxHeap.Values[maxI].Capital)) {
			maxI = rightC
		}
		if maxI != curI {
			maxHeap.Values[curI], maxHeap.Values[maxI] = maxHeap.Values[maxI], maxHeap.Values[curI]
			curI = maxI
		} else {
			break
		}
	}
}

func NewHeapWithVal() *MaxHeap {
	return &MaxHeap{
		Values:  make([]IPO, 0),
		CurSize: 0,
	}
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	aHeap := NewHeapWithVal()

	for i := 0; i < len(profits); i++ {
		anIPO := IPO{
			Profit:  profits[i],
			Capital: capital[i],
		}
		aHeap.AddElement(anIPO)
		//fmt.Printf("heap=%v\n",aHeap.Values)
	}
	//fmt.Printf("heap=%v\n",aHeap.Values)
	maxCap := w
	for i := 0; i < k; {
		if aHeap.CurSize > 0 && aHeap.Values[0].Capital <= maxCap {
			maxCap += aHeap.Values[0].Profit
			i++
			aHeap.RemoveElementTop()
			//fmt.Printf("After top removal: heap=%v\n", aHeap.Values)
		} else {
			leftI, curMaxInHeapLeft := findMaxProfitWithWealth(aHeap, maxCap, 1)
			rightI, curMaxInHeapRight := findMaxProfitWithWealth(aHeap, maxCap, 2)
			//fmt.Printf("i=%v curMaxInHeapLeft=%v curMaxInHeapRight=%v maxCap=%v\n",i,curMaxInHeapLeft,curMaxInHeapRight,maxCap)
			curMaxCap := int(math.Max(float64(curMaxInHeapLeft), float64(curMaxInHeapRight)))
			if curMaxCap != -1 {
				maxCap += curMaxCap
				i++
				if curMaxCap == curMaxInHeapLeft {
					aHeap.RemoveElementAtIndex(leftI)
				} else {
					aHeap.RemoveElementAtIndex(rightI)
				}
				//fmt.Printf("After Index removal: heap=%v\n", aHeap.Values)
			} else {
				break
			}
		}
	}
	return maxCap
}

func findMaxProfitWithWealth(maxHeap *MaxHeap, w, ind int) (int, int) {
	if ind >= len(maxHeap.Values) {
		return -1, -1
	}

	// main logic
	if maxHeap.Values[ind].Capital <= w {
		//return maxHeap.Values[ind].Profit
		return ind, maxHeap.Values[ind].Profit
	}
	leftI, leftM := findMaxProfitWithWealth(maxHeap, w, 2*ind+1)
	rightI, rightM := findMaxProfitWithWealth(maxHeap, w, 2*ind+2)
	if leftM > rightM {
		return leftI, leftM
	}
	return rightI, rightM
	//return int(math.Max(float64(leftM),float64(rightM)))
}
