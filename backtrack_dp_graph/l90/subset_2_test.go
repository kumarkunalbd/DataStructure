package l90_test

import (
	"kumarkunalbd/datastructure/backtrack_dp_graph/l90"
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	inputs := []struct {
		nums           []int
		expectedResult [][]int
	}{
		{
			[]int{4, 4, 4, 1, 4},
			[][]int{
				{1},
				{1, 4},
				{1, 4, 4},
				{1, 4, 4, 4},
				{1, 4, 4, 4, 4},
				{4},
				{4, 4},
				{4, 4, 4},
				{4, 4, 4, 4},
				{},
			},
		},
	}
	for _, item := range inputs {
		result := l90.SubsetsWithDup(item.nums)
		if reflect.DeepEqual(item.expectedResult, result) {
			t.Logf("SUCCESS, expected=%v  <-> got=%v\n", item.expectedResult, result)
		} else {
			t.Errorf("Failed,expected=%v  <-> got=%v\n", item.expectedResult, result)
		}
	}
}
