package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type SumRequest []int

func HandleSum(w http.ResponseWriter, r *http.Request) {
	slog.Info("/subtract hit")

	if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		slog.Error("Invalid method")
        return
    }

	var req SumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// Bad request
		http.Error(w, "Bad request", http.StatusBadRequest)
		slog.Error(fmt.Sprintf("Bad request: %v", r.Body))
		return
	}

	sum := 0
	for _, number := range req {
		sum += number
	}

	response := Response{
		Result: sum,
	}
	slog.Info(fmt.Sprintf("Computed sum of %v is %d", req, sum))

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}