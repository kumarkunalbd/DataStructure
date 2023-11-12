import (
	"sort"
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	numsC := make([]int, len(nums))
	_ = copy(numsC, nums)

	sort.Slice(numsC, func(i, j int) bool {
		return numsC[i] < numsC[j]
	})
	fmt.Printf("sorted=%v unsoretd=%#v\n", numsC, nums)

	s := -1
	e := -1
	for i, _ := range nums {
		if nums[i] != numsC[i] {
			s = i
			break
		}
	}

	if s != -1 {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] != numsC[i] {
				e = i
				break
			}
		}

		return e - s + 1
	}

	return 0
}