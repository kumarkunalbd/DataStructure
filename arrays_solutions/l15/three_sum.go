package l15

type SortedArr []int

func threeSum(nums []int) [][]int {
	sortedArr := SortedArr(nums)
	sortedArr.mergeSort(0, len(nums)-1)
	ans := make([][]int, 0)

	for i, val := range nums {
		if i > 0 && nums[i-1] == val {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			curSum := val + nums[l] + nums[r]
			if curSum < 0 {
				l++
			} else if curSum > 0 {
				r--
			} else {
				ans = append(ans, []int{val, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return ans
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
