package l78_test

import (
	"kumarkunalbd/datastructure/backtrack_dp_graph/l78"
	"reflect"
	"testing"
)

func TestSubsets3Input(t *testing.T) {
	input := []int{1, 2, 3}
	expectedResult := [][]int{
		[]int{},
		[]int{1},
		[]int{2},
		[]int{1, 2},
		[]int{3},
		[]int{1, 3},
		[]int{2, 3},
		[]int{1, 2, 3},
	}

	result := l78.Subsets(input)

	if reflect.DeepEqual(expectedResult, result) {
		t.Log("Test Passed with 3 input")
	} else {
		t.Errorf("Subsets Test failed. Expected Result:%v Actual result:%v\n", expectedResult, result)
	}
}
