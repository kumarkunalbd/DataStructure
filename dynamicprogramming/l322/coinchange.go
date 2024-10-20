package l322

import (
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	intil := []int{-1}
	memcoin := slices.Repeat(intil, amount+1)
	memcoin[0] = 0

	for i := 1; i < len(memcoin); i++ {
		var mincoin int
		mincoin = math.MaxInt
		if memcoin[i] == -1 {
			for _, val := range coins {
				if i-val >= 0 && memcoin[i-val] < mincoin {
					mincoin = memcoin[i-val]
				}
			}
		}
		if mincoin != math.MaxInt {
			mincoin++
		}
		memcoin[i] = mincoin
	}

	if memcoin[amount] == math.MaxInt {
		return -1
	}
	return memcoin[amount]
}
