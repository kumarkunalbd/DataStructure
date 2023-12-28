package l1522

type Node struct {
	Val      int
	Children []*Node
}

func diameter(root *Node) int {
	maxDia, _ := dfs(root)
	if maxDia > 0 {
		return maxDia - 1
	}
	return maxDia
}

func dfs(root *Node) (int, int) {
	// base case
	if root == nil {
		return 0, 0
	}
	// main logic
	maxDia := 0
	maxDepth := 0
	secondMaxDepth := 0

	for _, childNode := range root.Children {
		recDia, recDepth := dfs(childNode)
		if recDia > maxDia {
			maxDia = recDia
		}
		if recDepth > maxDepth {
			secondMaxDepth = maxDepth
			maxDepth = recDepth
		} else if recDepth > secondMaxDepth {
			secondMaxDepth = recDepth
		}
	}
	return maxValue(maxDia, maxDepth+secondMaxDepth+1), maxDepth + 1
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
