package BFS

func findRedundantConnection(edges [][]int) []int {
	maxN := 0
	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		maxN = max(maxN, max(x, y))
	}

	tolist := make([][]int, maxN+1)
	used := make([]bool, maxN+1)
	var dfs func(fa int, idx int)
	hascycle := false
	dfs = func(fa, idx int) {
		used[idx] = true
		for _, val := range tolist[idx] {
			if val == fa {
				continue
			}
			if !used[val] {
				dfs(idx, val)
			} else {
				hascycle = false
			}
		}
		used[idx] = false

	}
	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		tolist[x] = append(tolist[x], y)
		tolist[y] = append(tolist[y], x)
		dfs(0, 1)
		if hascycle {
			return edge
		}
	}
	return []int{}

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

type node struct {
	val int
	fa  int
}
