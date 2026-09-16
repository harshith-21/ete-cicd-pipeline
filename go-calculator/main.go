package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

type PageData struct {
	ThemeColor string
	Env 	   string
}

var hexColorRegex = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, "failed to read index.html", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("index.html").Parse(string(html))
	if err != nil {
		http.Error(w, "failed to parse index.html", http.StatusInternalServerError)
		return
	}

	themeColor := os.Getenv("THEME_COLOR")

	if !hexColorRegex.MatchString(themeColor) {
		themeColor = "#3498db"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = tmpl.Execute(w, PageData{
		ThemeColor: themeColor,
		Env:        os.Getenv("ENV"),
	})
	if err != nil {
		http.Error(w, "failed to render page", http.StatusInternalServerError)
		return
	}
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

	_ = json.NewEncoder(w).Encode(CalculateResponse{
		Result: result,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/calculate", calculateHandler)
	http.HandleFunc("/health", healthHandler)

	port := 8080

	log.Printf("Calculator running on http://localhost:%d", port)

	if err := http.ListenAndServe(
		":"+strconv.Itoa(port),
		nil,
	); err != nil {
		log.Fatal(err)
	}
}