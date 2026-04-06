package metrics

import "time"

type Execution struct {
	Solution []int `json:"solution"`
	Nodes    int   `json:"nodes"`
	TimeMs   int64 `json:"time_ms"`
}

func Measure(fn func() ([]int, int)) Execution {
	start := time.Now()

	solution, nodes := fn()

	return Execution{
		Solution: solution,
		Nodes:    nodes,
		TimeMs:   time.Since(start).Milliseconds(),
	}
}
