package algorithms

import "N-QUEENS_IN_GO/internal/domain/nqueens"

type Result struct {
	Solution []int
	Nodes    int
}

func DFS(n int) Result {
	nodes := 0

	var dfs func(state nqueens.State) ([]int, bool)

	dfs = func(state nqueens.State) ([]int, bool) {
		nodes++

		row := len(state)

		if row == n {
			return state, true
		}

		for col := 0; col < n; col++ {
			if nqueens.IsSafe(state, row, col) {
				newState := append(state, col)
				if sol, ok := dfs(newState); ok {
					return sol, true
				}
			}
		}

		return nil, false
	}

	solution, _ := dfs(nqueens.State{})
	return Result{Solution: solution, Nodes: nodes}
}
