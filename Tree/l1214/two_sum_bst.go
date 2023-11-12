package l1214

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	node_list1 := make([]int, 0)
	node_list1 = inOrder(root1, node_list1)
	node_list2 := make([]int, 0)
	node_list2 = inOrder(root2, node_list2)
	//fmt.Printf("list1=%v\n",node_list1)
	//fmt.Printf("list2=%v\n",node_list2)

	l := 0
	r := len(node_list2) - 1

	for l < len(node_list1) && r >= 0 {
		if node_list1[l]+node_list2[r] == target {
			return true
		} else if node_list1[l]+node_list2[r] < target {
			l++
		} else {
			r--
		}
	}
	return false
}

func inOrder(root *TreeNode, input []int) []int {
	// base case
	if root == nil {
		return input
	}

	// main logic
	input = inOrder(root.Left, input)
	input = append(input, root.Val)
	input = inOrder(root.Right, input)

	return input
}
