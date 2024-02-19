package l15

type SortedArr []int

func threeSum(nums []int) [][]int {
	sortedArr := SortedArr(nums)
	sortedArr.mergeSort(0, len(nums)-1)
	//fmt.Printf("Sorted:=%v\n",nums)
	ans := make([][]int, 0)
	l := 0
	for l < len(nums)-2 {
		curCombs := twoSum(nums, l+1, 0-nums[l])
		if len(curCombs) > 0 {
			ans = append(ans, curCombs...)
		}
		l++
		for l < len(nums)-2 && nums[l] == nums[l-1] {
			l++
		}
	}

	return ans
}

func twoSum(nums []int, sI int, target int) [][]int {
	l := sI
	r := len(nums) - 1
	curCombs := make([][]int, 0)
	for l < r {
		curSum := nums[l] + nums[r]
		if curSum == target {
			curCombs = append(curCombs, []int{nums[sI-1], nums[l], nums[r]})
			l++
			for l < len(nums) && nums[l] == nums[l-1] {
				l++
			}
		} else if curSum < target {
			l++
		} else {
			r--
		}
	}
	return curCombs
}

func (sortArr SortedArr) mergeSort(sI, eI int) {
	if sI >= eI {
		return
	}
	// main logic
	mid := (eI + sI) / 2
	sortArr.mergeSort(sI, mid)
	sortArr.mergeSort(mid+1, eI)
	sortArr.merge(sI, mid, mid+1, eI)
}

func (sortArr SortedArr) merge(s1, e1, s2, e2 int) {
	mergedArr := make([]int, e1-s1+1+e2-s2+1)

	l := s1
	r := s2
	counter := 0
	for l < s2 && r <= e2 {
		if sortArr[l] <= sortArr[r] {
			mergedArr[counter] = sortArr[l]
			l++
		} else {
			mergedArr[counter] = sortArr[r]
			r++
		}
		counter++
	}

	for l < s2 {
		mergedArr[counter] = sortArr[l]
		l++
		counter++
	}

	for r <= e2 {
		mergedArr[counter] = sortArr[r]
		r++
		counter++
	}

	for i, val := range mergedArr {
		sortArr[s1+i] = val
	}

}
