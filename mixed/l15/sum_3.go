package l15

import (
	"math"
	"math/rand"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sortInts(nums, 0, len(nums)-1)
	//fmt.Printf("array=%v\n", nums)
	prevNum := math.MinInt
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] != prevNum {
			newCombinations := getCombinations3Sum(nums, i)
			ans = append(ans, newCombinations...)
		}
		prevNum = nums[i]
	}

	return ans
}

func getCombinations3Sum(nums []int, ind int) [][]int {
	target := 0 - nums[ind]
	ans := make([][]int, 0)
	l := ind + 1
	r := len(nums) - 1
	mapFirstNum := make(map[int]bool, 0)
	for l < r {
		if nums[l]+nums[r] == target {
			if _, ok := mapFirstNum[nums[l]]; !ok {
				ans = append(ans, []int{nums[ind], nums[l], nums[r]})
				mapFirstNum[nums[l]] = true
			}
			l++
			r--
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			l++
		}
	}
	//fmt.Printf("ind=%v nums[ind]=%v add-ans=%v\n",ind,nums[ind],ans)
	return ans
}

func sortInts(nums []int, s, e int) {
	// base case
	if e <= s {
		return
	}

	// cmd logic
	randI := rand.Intn(e-s) + s
	nums[s], nums[randI] = nums[randI], nums[s]
	pN := nums[s]
	l := s + 1
	r := e
	for l <= r {
		if nums[l] <= pN {
			l++
		} else if nums[r] > pN {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	nums[r], nums[s] = nums[s], nums[r]

	sortInts(nums, s, r-1)
	sortInts(nums, r+1, e)
}
