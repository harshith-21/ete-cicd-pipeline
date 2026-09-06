package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CalculateRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	Op string  `json:"op"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalculateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var result float64

	switch req.Op {
	case "+":
		result = req.A + req.B
	case "-":
		result = req.A - req.B
	case "*":
		result = req.A * req.B
	case "/":
		if req.B == 0 {
			http.Error(w, "cannot divide by zero", http.StatusBadRequest)
			return
		}
		result = req.A / req.B
	default:
		http.Error(w, "invalid operator", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(CalculateResponse{
		Result: result,
	})
}

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)
	http.HandleFunc("/api/calculate", calculateHandler)

	port := 8080

	log.Printf("Calculator running on http://localhost:%d", port)

	if err := http.ListenAndServe(
		":"+strconv.Itoa(port),
		nil,
	); err != nil {
		log.Fatal(err)
	}
}