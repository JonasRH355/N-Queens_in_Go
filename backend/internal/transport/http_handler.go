package transport

import (
	"N-QUEENS_IN_GO/internal/algorithms"
	"encoding/json"
	"net/http"
	"time"
)

type Request struct {
	N         int    `json:"n"`
	Algorithm string `json:"algorithm"`
}

type Response struct {
	Solution []int `json:"solution"`
	Nodes    int   `json:"nodes"`
	TimeMs   int64 `json:"time_ms"`
}

func SolveHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	start := time.Now()

	var result algorithms.Result

	if req.Algorithm == "dfs" {
		result = algorithms.DFS(req.N)
	} else {
		result = algorithms.BFS(req.N)
	}

	resp := Response{
		Solution: result.Solution,
		Nodes:    result.Nodes,
		TimeMs:   time.Since(start).Milliseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
