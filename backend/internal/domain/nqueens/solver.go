package nqueens

func IsSafe(state State, row int, col int) bool {
	for r := 0; r < row; r++ {
		c := state[r]

		if c == col {
			return false
		}

		if abs(c-col) == abs(r-row) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
