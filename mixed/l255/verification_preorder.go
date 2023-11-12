package l255

import (
	"fmt"
	"math"
)

type Stack struct {
	Top    int
	Values []int
}

func VerifyPreorder(preorder []int) bool {
	aSt := NewStack()
	lastMinAllowed := math.MinInt

	for _, val := range preorder {
		//fmt.Printf("val=%v lastMinAllowed=%v\n",val,lastMinAllowed)
		if aSt.IsEmpty() {
			aSt.Push(val)
		} else {
			if val < aSt.Peek() {
				if val < lastMinAllowed {
					return false
				}
				aSt.Push(val)
			} else {
				for !aSt.IsEmpty() && val > aSt.Peek() {
					num, _ := aSt.Pop()
					lastMinAllowed = num
				}
				aSt.Push(val)
			}
		}
		//fmt.Printf("val=%v lastMinAllowed=%v stack=%v\n",val,lastMinAllowed,aSt)
	}

	return true

}

func NewStack() *Stack {
	return &Stack{
		Top:    -1,
		Values: make([]int, 0),
	}
}

func (st *Stack) Push(val int) {
	st.Values = append(st.Values, val)
	st.Top++
}

func (st *Stack) Pop() (int, error) {
	if st.Top == -1 {
		return math.MaxInt, fmt.Errorf("Stack is empty")
	}
	res := st.Values[len(st.Values)-1]
	st.Values = st.Values[0:(len(st.Values) - 1)]
	st.Top--
	return res, nil
}

func (st *Stack) Peek() int {
	return st.Values[len(st.Values)-1]
}

func (st *Stack) IsEmpty() bool {
	if st.Top == -1 {
		return true
	}

	return false
}
