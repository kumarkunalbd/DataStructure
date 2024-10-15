package l216

func CombinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	curL := make([]int, 0)
	ans = dfsComSum3(k, n, 1, curL, ans)
	return ans
}

func dfsComSum3(k, tarLeft, curNum int, curList []int, ans [][]int) [][]int {
	if len(curList) == k {
		return ans
	}
	// cmd logic
	i := curNum
	for {
		if i > 9 {
			break
		}
		newT := tarLeft - i
		if newT < 0 {
			break
		}
		if newT == 0 {
			if len(curList) == k-1 {
				newArr := make([]int, len(curList))
				copy(newArr, curList)
				newArr = append(newArr, i)
				ans = append(ans, newArr)
			}
			break
		} else if newT > 0 {
			curList = append(curList, i)
			ans = dfsComSum3(k, newT, i+1, curList, ans)
			curList = curList[:len(curList)-1]
		}
		i++
	}
	return ans
}
