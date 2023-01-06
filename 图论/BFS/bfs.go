package BFS

func canFinish(numCourses int, prerequisites [][]int) bool {
	tolist := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for _, val := range prerequisites {
		from := val[1]
		to := val[0]
		degree[to]++
		tolist[from] = append(tolist[from], to)
	}

	queue := make([]int, 0)
	for idx, val := range degree {
		if val == 0 {
			queue = append(queue, idx)
		}
	}
	lessons := 0
	for len(queue) != 0 {
		lessons++
		node := queue[0]
		queue = queue[1:]
		for _, val := range tolist[node] {
			degree[val]--
			if degree[val] == 0 {
				queue = append(queue, val)
			}
		}
	}
	if lessons != numCourses {
		return false
	}
	return true
}
