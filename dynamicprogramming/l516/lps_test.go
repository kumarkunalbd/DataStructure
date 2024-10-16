package l516

import (
	"log"
	"testing"
)

func TestLPS(t *testing.T) {
	inputs := []struct {
		text        string
		expectedres int
	}{
		{
			text:        "abkmba",
			expectedres: 5,
		},
		{
			text:        "akbma",
			expectedres: 3,
		},
		{
			"albkbrtap",
			5,
		},
	}

	for i, input := range inputs {
		actres := longestPalindromeSubseq(input.text)
		if input.expectedres != actres {
			t.Errorf("TestCase<%v> FAILED. Expcted:%v<->got:%v\n", i+1, input.expectedres, actres)
		} else {
			log.Printf("TestCAse<%v> PASSED.\n", i+1)
		}
	}
}
