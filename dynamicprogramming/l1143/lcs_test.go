package l1143

import (
	"log"
	"testing"
)

func TestLCS(t *testing.T) {
	inputs := []struct {
		text1       string
		text2       string
		expectedres int
	}{
		{
			text1:       "abc",
			text2:       "bck",
			expectedres: 2,
		},
		{
			text1:       "abcde",
			text2:       "akhbpkdjkeuj",
			expectedres: 4,
		},
		{
			text1:       "akhbpkdjkeuj",
			text2:       "abcdef",
			expectedres: 4,
		},
		{
			text1:       "akhbpkdjkeujabchjdefm",
			text2:       "abcdef",
			expectedres: 6,
		},
	}

	for i, input := range inputs {
		actualr := longestCommonSubsequence(input.text1, input.text2)
		if actualr != input.expectedres {
			t.Errorf("test Case <%v> FAILED. expected:%v<->got:%v\n", i, input.expectedres, actualr)
		} else {
			log.Printf("Test Case<%v> PASSED.", i)
		}
	}
}
