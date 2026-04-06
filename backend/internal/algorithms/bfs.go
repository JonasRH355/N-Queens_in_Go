package algorithms

import "N-QUEENS_IN_GO/internal/domain/nqueens"

func BFS(n int) Result {
	queue := []nqueens.State{{}}
	nodes := 0

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		nodes++

		row := len(state)

		if row == n {
			return Result{Solution: state, Nodes: nodes}
		}

		for col := 0; col < n; col++ {
			if nqueens.IsSafe(state, row, col) {
				newState := append(nqueens.State{}, state...)
				newState = append(newState, col)
				queue = append(queue, newState)
			}
		}
	}

	return Result{}
}
