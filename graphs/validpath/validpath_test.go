package validpath

import (
	"fmt"
	"log"
	"testing"
)

// Testvalidpath
func TestTheFunction(t *testing.T) {
	fmt.Println("---- Testvalidpath -----")
	type inputData struct {
		A int
		B int
		D int
		E []int
		F []int
	}
	inputs := []struct {
		indata         inputData
		expectedoutput string
	}{
		{
			indata: inputData{
				A: 2,
				B: 3,
				D: 1,
				E: []int{2},
				F: []int{3},
			},
			expectedoutput: "NO",
		},
		{
			indata: inputData{
				A: 3,
				B: 3,
				D: 1,
				E: []int{0},
				F: []int{3},
			},
			expectedoutput: "YES",
		},
	}

	for _, input := range inputs {
		actres := vpath(input.indata.A, input.indata.B, input.indata.D, input.indata.E, input.indata.F)
		if actres != input.expectedoutput {
			t.Errorf("Test Case failed : expected%v<->got:%v\n", input.expectedoutput, actres)
		} else {
			log.Printf("Test case passed")
		}
	}
}
