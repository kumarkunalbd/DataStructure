package l216_test

import (
	"kumarkunalbd/datastructure/backtrack_dp_graph/l216"
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {

	inputs := []struct {
		numsInComb     int
		tagetSum       int
		expectedResult [][]int
	}{
		{
			3,
			7,
			[][]int{
				{1, 2, 4},
			},
		},
		{
			3,
			9,
			[][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
		{
			4,
			1,
			[][]int{},
		},
	}

	for _, item := range inputs {
		result := l216.CombinationSum3(item.numsInComb, item.tagetSum)
		if reflect.DeepEqual(item.expectedResult, result) {
			t.Logf("PASSED, expected: %v <-> got :%v\n", item.expectedResult, result)
		} else {
			t.Errorf("FAILED, expected: %v <-> got :%v\n", item.expectedResult, result)
		}
	}
}
