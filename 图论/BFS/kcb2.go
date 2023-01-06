package BFS

func findOrder(numCourses int, prerequisites [][]int) []int {
	tolist := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for _, val := range prerequisites {
		from := val[1]
		to := val[0]
		degree[val[0]]++
		tolist[from] = append(tolist[from], to)
	}
	queue := make([]int, 0, numCourses)
	lessons := make([]int, 0, numCourses)
	for idx, val := range degree {
		if val == 0 {
			queue = append(queue, idx)
		}
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		lessons = append(lessons, node)
		for _, val := range tolist[node] {
			degree[val]--
			if degree[val] == 0 {
				queue = append(queue, val)
			}
		}
	}
	if len(lessons) != numCourses {
		return []int{}
	}
	return lessons
}
