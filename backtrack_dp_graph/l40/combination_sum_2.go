package l40

import "sort"

type sortedCan []int

func (list sortedCan) Len() int {
	return len(list)
}
func (list sortedCan) Less(i, j int) bool {
	return list[i] <= list[j]
}

func (list sortedCan) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func CombinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	curList := make([]int, 0)
	sort.Sort(sortedCan(candidates))
	//fmt.Printf("Cand=%v\n", candidates)
	ans = dfsCombinationSum2(candidates, target, 0, ans, curList)
	return ans
}

func dfsCombinationSum2(cand []int, target, curI int, ans [][]int, curList []int) [][]int {
	// main logic
	i := curI
	for {
		if i >= len(cand) {
			break
		}
		leftT := target - cand[i]
		if leftT < 0 {
			break
		}
		if leftT == 0 {
			newCopy := make([]int, len(curList))
			copy(newCopy, curList)
			newCopy = append(newCopy, cand[i])
			ans = append(ans, newCopy)
			break
		} else if leftT > 0 {
			curList = append(curList, cand[i])
			ans = dfsCombinationSum2(cand, leftT, i+1, ans, curList)
			curList = curList[:len(curList)-1]
			for {
				i++
				if i == len(cand) || cand[i] != cand[i-1] {
					break
				}
			}
		}
	}
	return ans
}
