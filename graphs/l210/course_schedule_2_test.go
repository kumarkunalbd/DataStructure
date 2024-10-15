package l210

import (
	"log"
	"reflect"
	"testing"
)

func TestCourseSchedule2(t *testing.T) {
	type inputpatter struct {
		prereqs    [][]int
		numcourses int
	}
	datainputs := []struct {
		actualinput inputpatter
		expectedres []int
	}{
		{
			actualinput: inputpatter{
				prereqs:    [][]int{[]int{1, 0}},
				numcourses: 2,
			},
			expectedres: []int{0, 1},
		},
		{
			actualinput: inputpatter{
				prereqs: [][]int{
					[]int{1, 0},
					[]int{2, 0},
					[]int{3, 1},
					[]int{3, 2},
				},
				numcourses: 4,
			},
			expectedres: []int{0, 1, 2, 3},
		},
		{
			actualinput: inputpatter{
				prereqs: [][]int{
					[]int{0, 1},
					[]int{1, 0},
				},
				numcourses: 2,
			},
			expectedres: []int{},
		},
	}

	for i, input := range datainputs {
		actres := findOrder(input.actualinput.numcourses, input.actualinput.prereqs)
		if !reflect.DeepEqual(actres, input.expectedres) {
			t.Errorf("Testcase <%v> FAILED. Exptected:%v<->got:%v\n", i, input.expectedres, actres)
		} else {
			log.Printf("Testcase <%v> SUCCEED.\n", i)
		}
	}
}
