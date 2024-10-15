package l207

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjl := getAdjList(prerequisites)
	//fmt.Printf("adjl=%v\n",adjl)
	visited := make(map[int]bool, 0)
	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			//fmt.Printf("finding i=%v startv=%v\n",i,visited)
			visited[i] = true
			lineage := make(map[int]bool, 0)
			lineage[i] = true
			if isCyclic(adjl, i, visited, lineage) {
				return false
			}
			//fmt.Printf("endvisted=%v\n",visited)
		}
	}
	return true
}

func isCyclic(adjl map[int]map[int]bool, curv int, visited map[int]bool, lineage map[int]bool) bool {

	// main logic
	for preqv := range adjl[curv] {
		if lineage[preqv] {
			return true
		}
		if !visited[preqv] {
			visited[preqv] = true

			lineage[preqv] = true
			if isCyclic(adjl, preqv, visited, lineage) {
				return true
			}
			delete(lineage, preqv)
		}
	}
	return false
}

func getAdjList(preqs [][]int) map[int]map[int]bool {
	adjl := make(map[int]map[int]bool)

	for _, preq := range preqs {
		if adjl[preq[0]] == nil {
			adjl[preq[0]] = make(map[int]bool)
		}
		adjl[preq[0]][preq[1]] = true
	}

	return adjl
}
