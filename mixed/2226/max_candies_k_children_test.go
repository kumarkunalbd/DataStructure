package leet2226

import (
	//"fmt"
	"fmt"
	"testing"
)

type addTest struct {
	arg1     []int
	arg2     int64
	expected int
}

var addTests = []addTest{
	addTest{[]int{5, 8, 6}, 3, 5},
	addTest{[]int{2, 5}, 11, 0},
	addTest{[]int{2, 5, 7}, 11, 1},
	addTest{[]int{2, 5, 8}, 11, 1},
}

func TestMaximumCandies(t *testing.T) {
	for _, test := range addTests {
		if output := maximumCandies(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %v not equal to expected %v\n", output, test.expected)
		} else {
			t.Log(fmt.Printf("Output %v equal to expected %v\n", output, test.expected))
		}
	}
}

func BenchmarkMaximumCandies(b *testing.B) {
	//test := addTests[0]
	for _, test := range addTests {
		for i := 0; i < b.N; i++ {
			maximumCandies(test.arg1, test.arg2)
		}
	}
}
