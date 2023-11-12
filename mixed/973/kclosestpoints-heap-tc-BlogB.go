import (
	"fmt"
)

func kClosest(points [][]int, k int) [][]int {
	maxHeap := make([][]int, 0)
	for i := 0; i < k; i++ {
		maxHeap = addElementMaxHeap(points[i], maxHeap)
	}
	//fmt.Printf("l1: maxHeap=%v\n",maxHeap)
	for i := k; i < len(points); i++ {
		maxHeapED := (maxHeap[0][0] * maxHeap[0][0]) + (maxHeap[0][1] * maxHeap[0][1])
		curED := (points[i][0] * points[i][0]) + (points[i][1] * points[i][1])

		if curED < maxHeapED {
			maxHeap = deleteElementMaxHeap(maxHeap)
			maxHeap = addElementMaxHeap(points[i], maxHeap)
		}
	}
	//fmt.Printf("l2: maxHeap=%v\n",maxHeap)
	return maxHeap
}

func addElementMaxHeap(point []int, maxHeap [][]int) [][]int {
	maxHeap = append(maxHeap, point)
	ptr := len(maxHeap) - 1
	//fmt.Printf("addElementMaxHeap1: ptr=%v maxHeap=%v\n",ptr,maxHeap)
	for ptr > 0 {
		parent := (ptr - 1) / 2
		ptrED := (maxHeap[ptr][0] * maxHeap[ptr][0]) + (maxHeap[ptr][1] * maxHeap[ptr][1])
		parentED := (maxHeap[parent][0] * maxHeap[parent][0]) + (maxHeap[parent][1] * maxHeap[parent][1])
		if parentED < ptrED {
			maxHeap[parent], maxHeap[ptr] = maxHeap[ptr], maxHeap[parent]
			//fmt.Printf("addElementMaxHeap2: ptr=%v ptrED=%v parent=%v parentED=%v maxHeap=%v\n",ptr,ptrED,parent,parentED,maxHeap)
		} else {
			break
		}
		ptr = parent

	}

	return maxHeap
}

func deleteElementMaxHeap(maxHeap [][]int) [][]int {
	maxHeap[0], maxHeap[len(maxHeap)-1] = maxHeap[len(maxHeap)-1], maxHeap[0]

	maxHeap = maxHeap[:len(maxHeap)-1]
	ptr := 0
	for ptr < len(maxHeap) {
		leftC := 2*ptr + 1
		if leftC >= len(maxHeap) {
			break
		}
		maxEle := ptr
		maxEleED := (maxHeap[maxEle][0] * maxHeap[maxEle][0]) + (maxHeap[maxEle][1] * maxHeap[maxEle][1])
		leftCED := (maxHeap[leftC][0] * maxHeap[leftC][0]) + (maxHeap[leftC][1] * maxHeap[leftC][1])
		if leftCED > maxEleED {
			maxEle = leftC
			maxEleED = leftCED
		}
		rightC := 2*ptr + 2
		if rightC >= len(maxHeap) {
			if maxEle != ptr {
				maxHeap[ptr], maxHeap[maxEle] = maxHeap[maxEle], maxHeap[ptr]
			}
			break
		}
		rightCED := maxHeap[rightC][0]*maxHeap[rightC][0] + maxHeap[rightC][1]*maxHeap[rightC][1]
		if rightCED > maxEleED {
			maxEle = rightC
		}

		if maxEle != ptr {
			maxHeap[ptr], maxHeap[maxEle] = maxHeap[maxEle], maxHeap[ptr]
			ptr = maxEle
		} else {
			break
		}
	}

	return maxHeap

}