package d37

func solveSudoku(board [][]byte) {
	rows := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 10)
	}
	cols := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		cols[i] = make([]bool, 10)
	}
	boxes := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		boxes[i] = make([]bool, 10)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			digit := board[i][j] - '0'
			rows[i][digit] = true
			cols[j][digit] = true
			boxes[(i/3)*3+(j/3)][digit] = true
		}
	}
	GetLen := func(i, j int) []int {
		digits := make([]int, 0, 10)
		for k := 1; k < 9; k++ {
			if rows[i][k] && cols[j][k] && boxes[(i/3)*3+j/3][k] {
				digits = append(digits, k)
			}
		}
		return digits
	}
	GetBestPos := func() []int {
		min := 10
		minPos := []int{-1, -1}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				digits := GetLen(i, j)
				if min > len(digits) {
					min = len(digits)
					minPos[0] = i
					minPos[1] = j
				}
			}
		}
		return minPos
	}
	var dfs func() bool
	dfs = func() bool {
		pos := GetBestPos()
		i := pos[0]
		j := pos[1]
		if i == -1 {
			return true
		}
		digits := GetLen(pos[0], pos[1])

		for _, digit := range digits {
			rows[i][digit] = true
			cols[j][digit] = true
			boxes[(i/3)*3+(j/3)][digit] = true
			board[i][j] = byte(digit + '0')
			if dfs() {
				return true
			}
			board[i][j] = '.'
			boxes[(i/3)*3+(j/3)][digit] = false
			cols[j][digit] = false
			rows[i][digit] = false

		}
		return false
	}

}
