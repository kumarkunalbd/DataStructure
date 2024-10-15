package l502

import "math/rand"

type MaxHeap struct {
	Values []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		Values: make([]int, 0),
	}
}

func (maxHeap *MaxHeap) AddElement(val int) {
	maxHeap.Values = append(maxHeap.Values, val)
	curI := len(maxHeap.Values) - 1
	for curI > 0 {
		parentI := (curI - 1) / 2
		if maxHeap.Values[parentI] < maxHeap.Values[curI] {
			maxHeap.Values[curI], maxHeap.Values[parentI] = maxHeap.Values[parentI], maxHeap.Values[curI]
			curI = parentI
		} else {
			break
		}
	}
}

func (maxHeap *MaxHeap) RemoveTopElement() {
	maxHeap.Values[0], maxHeap.Values[len(maxHeap.Values)-1] = maxHeap.Values[len(maxHeap.Values)-1], maxHeap.Values[0]
	maxHeap.Values = maxHeap.Values[:len(maxHeap.Values)-1]

	curI := 0

	for curI < len(maxHeap.Values) {
		maxI := curI
		leftC := curI*2 + 1
		if leftC < len(maxHeap.Values) && maxHeap.Values[leftC] > maxHeap.Values[maxI] {
			maxI = leftC
		}
		rightC := curI*2 + 2
		if rightC < len(maxHeap.Values) && maxHeap.Values[rightC] > maxHeap.Values[maxI] {
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

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	aHeap := NewMaxHeap()
	quickSortTheArrays(profits, capital, 0, len(profits)-1)
	//fmt.Printf("capital=%v\n",capital)
	//fmt.Printf("profits=%v\n",profits)
	notAvailablePtr := 0
	maxC := w
	for notAvailablePtr < len(capital) {
		if capital[notAvailablePtr] <= w {
			aHeap.AddElement(profits[notAvailablePtr])
			notAvailablePtr++
		} else {
			break
		}
	}

	//fmt.Printf("Heap=%v\n",aHeap)

	for i := 0; i < k && len(aHeap.Values) != 0; i++ {
		maxC += aHeap.Values[0]
		aHeap.RemoveTopElement()
		for notAvailablePtr < len(capital) {
			if capital[notAvailablePtr] <= maxC {
				aHeap.AddElement(profits[notAvailablePtr])
				notAvailablePtr++
			} else {
				break
			}
		}
	}
	return maxC
}

func quickSortTheArrays(profits, capital []int, startI, endI int) {
	// base case
	if startI >= endI {
		return
	}

	// cmd logic
	partI := partition(profits, capital, startI, endI)
	quickSortTheArrays(profits, capital, startI, partI-1)
	quickSortTheArrays(profits, capital, partI+1, endI)
}

func partition(profits, capital []int, startI, endI int) int {
	//fmt.Printf("Partition profit=%v capital=%v startI=%v endI=%v\n",profits,capital,startI,endI)
	pivot := rand.Intn(endI-startI) + startI
	capital[pivot], capital[startI] = capital[startI], capital[pivot]
	profits[pivot], profits[startI] = profits[startI], profits[pivot]
	pivot = startI
	l := startI + 1
	r := endI

	for l <= r {
		if capital[l] <= capital[pivot] {
			l++
		} else if capital[r] > capital[pivot] {
			r--
		} else {
			capital[l], capital[r] = capital[r], capital[l]
			profits[l], profits[r] = profits[r], profits[l]
			l++
			r--
		}
	}
	capital[pivot], capital[r] = capital[r], capital[pivot]
	profits[pivot], profits[r] = profits[r], profits[pivot]
	return r
}
