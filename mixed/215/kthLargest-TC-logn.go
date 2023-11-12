func findKthLargest(nums []int, k int) int {
	minHeap := make([]int, 0)
	for i := 0; i < k; i++ {
		minHeap = addElement(minHeap, nums[i])
	}

	//fmt.Printf("L1: minHeap=%v\n",minHeap)
	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap[0] {
			minHeap = removeElement(minHeap)
			minHeap = addElement(minHeap, nums[i])
		}

		//fmt.Printf("L2: i=%d minHerap=%v\n",i,minHeap)
	}
	return minHeap[0]
}

func addElement(minHeap []int, num int) []int {
	minHeap = append(minHeap, num)
	curI := len(minHeap) - 1
	for curI > 0 {
		parentI := (curI - 1) / 2
		if minHeap[curI] < minHeap[parentI] {
			minHeap[curI], minHeap[parentI] = minHeap[parentI], minHeap[curI]
		} else {
			break
		}
		curI = parentI
	}
	return minHeap
}

func removeElement(minHeap []int) []int {
	minHeap[len(minHeap)-1], minHeap[0] = minHeap[0], minHeap[len(minHeap)-1]
	minHeap = minHeap[:len(minHeap)-1]

	curI := 0

	for curI < len(minHeap) {
		leftC := 2*curI + 1
		rightC := 2*curI + 2
		minI := curI

		if leftC < len(minHeap) && minHeap[leftC] < minHeap[minI] {
			minI = leftC
		}
		if rightC < len(minHeap) && minHeap[rightC] < minHeap[minI] {
			minI = rightC
		}

		if minI == curI {
			break
		} else {
			minHeap[curI], minHeap[minI] = minHeap[minI], minHeap[curI]
		}
		curI = minI
	}
	return minHeap
}