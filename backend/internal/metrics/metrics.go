// package metrics

// import "time"

// type Execution struct {
// 	Solution []int `json:"solution"`
// 	Nodes    int   `json:"nodes"`
// 	TimeMs   int64 `json:"time_ms"`
// }

// func Measure(fn func() ([]int, int)) Execution {
// 	start := time.Now()

// 	solution, nodes := fn()

// 	return Execution{
// 		Solution: solution,
// 		Nodes:    nodes,
// 		TimeMs:   time.Since(start).Milliseconds(),
// 	}
// }

package metrics

import (
	"runtime"
	"time"
)

type Execution struct {
	Solution []int  `json:"solution"`
	Nodes    int    `json:"nodes"`
	TimeMs   int64  `json:"time_ms"`
	Memory   uint64 `json:"memory"`
}

func Measure(fn func() ([]int, int)) Execution {
	var mStart, mEnd runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&mStart)

	start := time.Now()

	solution, nodes := fn()

	elapsed := time.Since(start)

	runtime.ReadMemStats(&mEnd)

	usedMemory := (mEnd.Alloc - mStart.Alloc)

	return Execution{
		Solution: solution,
		Nodes:    nodes,
		TimeMs:   elapsed.Milliseconds(),
		Memory:   usedMemory,
	}
}
