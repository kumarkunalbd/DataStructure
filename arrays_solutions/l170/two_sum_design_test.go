package l170_test

import (
	"kumarkunalbd/datastructure/arrays_solutions/l170"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	trueV := true
	falseV := false
	inputs := []struct {
		instructions []string
		numbersArr   [][]int
		results      []*bool
	}{
		{
			[]string{"TwoSum", "add", "add", "add", "find", "find"},
			[][]int{
				{},
				{1},
				{3},
				{5},
				{4},
				{7},
			},
			[]*bool{nil, nil, nil, nil, &trueV, &falseV},
		},
		{
			[]string{"TwoSum", "add", "add", "add", "find"},
			[][]int{
				{},
				{0},
				{-1},
				{1},
				{0},
			},
			[]*bool{nil, nil, nil, nil, &trueV},
		},
	}
	var obj *l170.TwoSum
	for _, input := range inputs {
		for i, instruction := range input.instructions {
			switch instruction {
			case "TwoSum":
				twS := l170.Constructor()
				obj = &twS
				if obj == nil {
					t.Errorf("Failed,expected=%v  <-> got=%v\n", *obj, nil)
				} else {
					t.Logf("SUCCESS, expected=%v  <-> got=%v\n", *obj, *obj)
				}
			case "add":
				if obj == nil {
					panic("Error: Pass TwoSum instrcutions in iput array")
					return
				}
				numArr := input.numbersArr[i]
				if len(numArr) == 0 {
					panic("Error: Pass a num in input of numbers")
					return
				}
				obj.Add(numArr[0])
				if len(obj.MapFreq) == 0 {
					t.Errorf("Failed,expected lenth=%v  <-> got=%v\n", ">0", "0")
				} else {
					t.Logf("SUCCESS, expected length=%v  <-> got=%v\n", len(obj.MapFreq), len(obj.MapFreq))
				}
			case "find":
				numArr := input.numbersArr[i]
				actRes := obj.Find(numArr[0])
				if actRes != *input.results[i] {
					t.Errorf("Failed,expected =%v  <-> got=%v\n", *input.results[i], actRes)
				} else {
					t.Logf("SUCCESS, expected length=%v  <-> got=%v\n", *input.results[i], actRes)
				}
			default:

			}
		}
	}
}
