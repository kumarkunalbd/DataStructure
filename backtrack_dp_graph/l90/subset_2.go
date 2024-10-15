package l90

import "math/rand"

type custSlc []int

func (anArr custSlc) sortAsc(startI, endI int) {
	//fmt.Printf("startI=%v endI=%v\n",startI,endI)
	if startI >= endI {
		return
	}
	// cmd logic
	randNum := rand.Intn(endI-startI) + startI
	anArr[randNum], anArr[startI] = anArr[startI], anArr[randNum]
	pivot := startI
	l := startI + 1
	r := endI
	for l <= r {
		if anArr[l] <= anArr[pivot] {
			l++
		} else if anArr[r] > anArr[pivot] {
			r--
		} else {
			anArr[l], anArr[r] = anArr[r], anArr[l]
			l++
			r--
		}
	}
	anArr[pivot], anArr[r] = anArr[r], anArr[pivot]
	//fmt.Printf("arr=%v randNum=%v r=%v\n",anArr,randNum,r)
	anArr.sortAsc(startI, r-1)
	anArr.sortAsc(r+1, endI)
}

func SubsetsWithDup(nums []int) [][]int {
	aCustArr := custSlc(nums)
	aCustArr.sortAsc(0, len(nums)-1)
	//fmt.Printf("sorted arr=%v\n",[]int(aCustArr))
	ans := make([][]int, 0)
	curArr := make([]int, 0)
	ans = dfsSubsetsWithDup([]int(aCustArr), curArr, 0, ans)
	if len(ans) > 0 {
		ans = append(ans, make([]int, 0))
	}
	return ans
}

func dfsSubsetsWithDup(nums []int, curArr []int, startInd int, ans [][]int) [][]int {

	// cmd logic
	for i := startInd; i < len(nums); i++ {
		newArr := append(curArr, nums[i])
		finalArr := make([]int, len(newArr))
		copy(finalArr, newArr)
		ans = append(ans, finalArr)
		ans = dfsSubsetsWithDup(nums, newArr, i+1, ans)
		for i+1 < len(nums) && nums[i+1] == nums[i] {
			i++
		}
	}
	return ans
}
