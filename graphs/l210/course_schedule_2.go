package l210

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjl := getAdjList(prerequisites)
	//fmt.Printf("adjl=%v\n",adjl)
	finishOrder := make([]int, 0)
	visited := make(map[int]bool, 0)
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			visited[i] = true
			curlineage := make(map[int]bool, 0)
			curlineage[i] = true
			if isCyclic(adjl, i, visited, curlineage, &finishOrder) {
				return []int{}
			}
		}
	}
	return finishOrder
}

func getAdjList(preqs [][]int) map[int]map[int]bool {
	adjl := make(map[int]map[int]bool, 0)

	for _, preq := range preqs {
		if adjl[preq[0]] == nil {
			adjl[preq[0]] = make(map[int]bool, 0)
		}
		adjl[preq[0]][preq[1]] = true
	}
	return adjl
}

func isCyclic(adjl map[int]map[int]bool, curv int, visited map[int]bool, curlin map[int]bool, finord *[]int) bool {
	// main logic
	for nv := range adjl[curv] {
		if curlin[nv] {
			return true
		}
		if !visited[nv] {
			visited[nv] = true
			curlin[nv] = true
			if isCyclic(adjl, nv, visited, curlin, finord) {
				return true
			}
			delete(curlin, nv)
		}
	}
	*finord = append(*finord, curv)
	return false
}
