func maxDistance(position []int, m int) int {
	position = mergeSortInts(position, 0, len(position)-1)
	//fmt.Printf("Sorted Array=%v\n",position)
	s := 1
	e := position[len(position)-1] - position[0]
	largestMin := -1
	for s <= e {
		mid := (s + e) / 2
		//fmt.Printf("l1: mid=%v s=%v e=%v largestMin=%v\n",mid,s,e,largestMin)
		if sizeFit(position, mid, m) {
			largestMin = mid
			s = mid + 1
		} else {
			e = mid - 1
		}
		//fmt.Printf("l2: mid=%v s=%v e=%v largestMin=%v\n",mid,s,e,largestMin)
	}
	return largestMin
}

func sizeFit(nums []int, size, B int) bool {
	curMinStallReq := nums[0]
	for _, val := range nums {
		if val >= curMinStallReq {
			B--
			curMinStallReq = val + size
		}
		//fmt.Printf("B=%v curMinReq=%v\n",B,curMinStallReq)
	}
	if B <= 0 {
		return true
	}
	return false
}

func mergeSortInts(nums []int, s, e int) []int {
	// base case
	if s >= e {
		return nums
	}
	// main logic
	mid := s + (e-s+1)/2
	nums = mergeSortInts(nums, s, mid-1)
	nums = mergeSortInts(nums, mid, e)
	nums = mergeNums(nums, s, mid, e)
	return nums
}

func mergeNums(nums []int, s1, s2, e2 int) []int {
	supArr := make([]int, e2-s1+1)
	l := s1
	r := s2
	count := 0
	for l < s2 && r <= e2 {
		if nums[l] <= nums[r] {
			supArr[count] = nums[l]
			l++
		} else {
			supArr[count] = nums[r]
			r++
		}
		count++
	}

	for l < s2 {
		supArr[count] = nums[l]
		l++
		count++
	}

	for r <= e2 {
		supArr[count] = nums[r]
		r++
		count++
	}

	for i, val := range supArr {
		nums[s1+i] = val
	}
	//fmt.Printf("merged=%v s1=%v s2=%v e2=%v\n",nums,s1,s2,e2)

	return nums
}