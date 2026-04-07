package transport

import (
	"N-QUEENS_IN_GO/internal/algorithms"
	"N-QUEENS_IN_GO/internal/metrics"
	"encoding/json"
	"net/http"
	"time"
)

type Request struct {
	N         int    `json:"n"`
	Algorithm string `json:"algorithm"`
}

type Response struct {
	Solution []int  `json:"solution"`
	Nodes    int    `json:"nodes"`
	TimeMs   int64  `json:"time_ms"`
	Memory   uint64 `json:"memory"`
}

func SolveHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()

	// var result algorithms.Result

	// if req.Algorithm == "dfs" {
	// 	result = algorithms.DFS(req.N)
	// } else {
	// 	result = algorithms.BFS(req.N)
	// }

	result := metrics.Measure(func() ([]int, int) {
		if req.Algorithm == "dfs" {
			r := algorithms.DFS(req.N)
			return r.Solution, r.Nodes
		} else {
			r := algorithms.BFS(req.N)
			return r.Solution, r.Nodes
		}
	})

	resp := Response{
		Solution: result.Solution,
		Nodes:    result.Nodes,
		TimeMs:   time.Since(start).Milliseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
