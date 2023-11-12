package leet1870

import (
	"fmt"
	"testing"
)

type testCase struct {
	arg1     []int
	arg2     float64
	expected int
}

var minSpeedTestCases = []testCase{
	testCase{[]int{1, 3, 2}, 6, 1},
	testCase{[]int{1, 3, 2}, 2.7, 3},
	testCase{[]int{1, 3, 2}, 1.9, -1},
	testCase{[]int{69}, 4.6, 15},
}

func TestMinSpeedOnTime(t *testing.T) {
	for _, test := range minSpeedTestCases {
		if output := minSpeedOnTime(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output=%v is not equal to expected=%v\n", output, test.expected)
		} else {
			t.Log(fmt.Printf("Output=%v is equal to expected=%v\n", output, test.expected))
		}
	}
}
