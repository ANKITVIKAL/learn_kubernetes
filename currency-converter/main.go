package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const frankfurterBase = "https://api.frankfurter.app"

func main() {
	http.HandleFunc("/api/currencies", handleCurrencies)
	http.HandleFunc("/api/latest", handleLatest)
	http.HandleFunc("/", handleIndex)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Currency converter server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleCurrencies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := http.Get(frankfurterBase + "/currencies")
	if err != nil {
		http.Error(w, "failed to fetch currencies", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("write currencies: %v", err)
	}
}

func handleLatest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amount := r.URL.Query().Get("amount")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	if amount == "" || from == "" || to == "" {
		http.Error(w, "missing amount, from, or to", http.StatusBadRequest)
		return
	}

	url := frankfurterBase + "/latest?amount=" + amount + "&from=" + from + "&to=" + to
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "failed to fetch rate", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "failed to read response", http.StatusInternalServerError)
		return
	}

	// Pass through Frankfurter response as-is so frontend logic stays the same
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "invalid response from rates API", http.StatusBadGateway)
		return
	}
	if msg, ok := data["message"].(string); ok && msg != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	path := filepath.Join(".", "index.html")
	http.ServeFile(w, r, path)
}
