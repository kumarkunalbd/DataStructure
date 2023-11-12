/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */
package l1265

import "fmt"

type stack struct {
	top        int
	storageArr []ImmutableListNode
}

func (thisS *stack) pop() (ImmutableListNode, error) {
	if thisS.size() > 0 {
		//fmt.Printf("l1:Arra is:%v\n",thisS.storageArr)
		val := thisS.storageArr[thisS.top]
		thisS.storageArr = thisS.storageArr[:(len(thisS.storageArr) - 1)]
		//fmt.Printf("l2:Arra is:%v\n",thisS.storageArr)
		thisS.top--
		return val, nil
	}
	return nil, fmt.Errorf("Stack empty")

}

func (thisS *stack) seek() ImmutableListNode {
	return thisS.storageArr[thisS.top]
}

func (thisS *stack) push(val ImmutableListNode) {
	thisS.storageArr = append(thisS.storageArr, val)
	thisS.top++
}

func (thisS *stack) size() int {
	return len(thisS.storageArr)
}

func NewStackWithVal(val ImmutableListNode) *stack {
	return &stack{
		top:        0,
		storageArr: make([]ImmutableListNode, 0),
	}
}

func NewStack() *stack {
	return &stack{
		top:        -1,
		storageArr: make([]ImmutableListNode, 0),
	}
}

func printLinkedListInReverse(head ImmutableListNode) {
	ns := NewStack()
	for head != nil {
		ns.push(head)
		head = head.getNext()
	}

	for ns.size() > 0 {
		fmt.Printf("size:%v top:%v\n", ns.size(), ns.top)
		posItem, _ := ns.pop()
		posItem.printValue()
	}
}
