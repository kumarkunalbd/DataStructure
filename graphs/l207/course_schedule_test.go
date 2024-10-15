package l207

import (
	"log"
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	type inputtemplate struct {
		prereqs    [][]int
		numcourses int
	}

	inputs := []struct {
		in          inputtemplate
		expectedres bool
	}{
		{
			in: inputtemplate{
				prereqs:    [][]int{{1, 0}},
				numcourses: 2,
			},
			expectedres: true,
		},
		{
			in: inputtemplate{
				prereqs:    [][]int{[]int{1, 0}, []int{0, 1}},
				numcourses: 2,
			},
			expectedres: false,
		},
		{
			in: inputtemplate{
				prereqs:    [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{2, 5}, []int{1, 4}, []int{4, 5}},
				numcourses: 6,
			},
			expectedres: true,
		},
	}

	for i, input := range inputs {
		actres := canFinish(input.in.numcourses, input.in.prereqs)
		if actres != input.expectedres {
			t.Errorf("Test Failed : expected:%v<->got:%v\n", input.expectedres, actres)
		} else {
			log.Printf("Test Case <%v> passed\n", i+1)
		}
	}
}
