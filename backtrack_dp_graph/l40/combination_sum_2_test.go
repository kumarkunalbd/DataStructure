package l40_test

import (
	"kumarkunalbd/datastructure/backtrack_dp_graph/l40"
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	inputs := []struct {
		candidates     []int
		target         int
		expectedResult [][]int
	}{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
	}
	for _, item := range inputs {
		result := l40.CombinationSum2(item.candidates, item.target)
		if reflect.DeepEqual(item.expectedResult, result) {
			t.Logf("SUCCESS, expected=%v  <-> got=%v\n", item.expectedResult, result)
		} else {
			t.Errorf("Failed,expected=%v  <-> got=%v\n", item.expectedResult, result)
		}
	}
}
