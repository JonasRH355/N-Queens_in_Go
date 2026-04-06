package main

import (
	"N-QUEENS_IN_GO/internal/transport"
	"net/http"
)

func main() {
	http.HandleFunc("/solve", transport.SolveHandler)

	println("API rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
