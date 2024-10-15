package l206

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	mappos := make(map[int]int, 0)
	for i, num := range inorder {
		mappos[num] = i
	}
	root := buildTreeRec(inorder, postorder, mappos, 0, len(inorder)-1, len(postorder)-1)
	//PrintT(root)
	return root
}

func buildTreeRec(inorder []int, postorder []int, mappos map[int]int, sin int, ein int, postrooti int) *TreeNode {
	//fmt.Printf("sin=%v ein=%v postrooti=%v\n",sin,ein,postrooti)
	// base cases
	if postrooti < 0 || ein < sin {
		//fmt.Printf("Back : nil\n")
		return nil
	}
	if sin == ein {

		cr := &TreeNode{
			Val:   postorder[postrooti],
			Left:  nil,
			Right: nil,
		}
		//fmt.Printf("Back : cr=%v\n",*cr)
		return cr
	}

	// cmd logic
	cr := &TreeNode{
		Val:   postorder[postrooti],
		Left:  nil,
		Right: nil,
	}
	posin := mappos[postorder[postrooti]]
	//fmt.Printf("L1: posin=%v cr=%v sin=%v postrooti=%v\n",posin,*cr,sin,postrooti)
	rt := buildTreeRec(inorder, postorder, mappos, posin+1, ein, postrooti-1)
	//fmt.Printf("rt=%v\n",rt)
	cr.Right = rt
	rtsize := ein - posin
	//fmt.Printf("L2: posin=%v cr=%v sin=%v postrooti=%v rtsize=%v\n",posin,*cr,sin,postrooti,rtsize)
	lt := buildTreeRec(inorder, postorder, mappos, sin, posin-1, postrooti-rtsize-1)
	//fmt.Printf("lt=%v\n",lt)
	cr.Left = lt
	return cr
}

func PrintT(root *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		dqn := queue[0]
		if dqn != nil {
			fmt.Printf("%v ", *dqn)
			queue = append(queue, dqn.Left)
			queue = append(queue, dqn.Right)
		} else {
			fmt.Printf("null ")
		}
		queue = queue[1:]
	}

	fmt.Println("END")
}
