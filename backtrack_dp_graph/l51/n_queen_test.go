package l51_test

import (
	"kumarkunalbd/datastructure/backtrack_dp_graph/l51"
	"reflect"
	"testing"
)

func TestDfsNQueen(t *testing.T) {
	inputs := []struct {
		n      int
		result [][]string
	}{
		{
			4,
			[][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			1,
			[][]string{
				{"Q"},
			},
		},
	}
	for _, item := range inputs {
		actualResult := l51.SolveNQueens(item.n)
		if reflect.DeepEqual(actualResult, item.result) {
			t.Logf("TestDfsNQueen SUCCCEED. Expected->%v  got->%v\n", item.result, actualResult)
		} else {
			t.Errorf("TestDFCNQueen FAILED. Expected->%v got->%v\n", item.result, actualResult)
		}
	}
}
